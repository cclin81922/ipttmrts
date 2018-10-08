package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Station ...
type Station struct {
	NameTW    string
	NameEN    string
	Latitude  float64
	Longitude float64
	Distance  float64
}

func (s Station) String() string {
	return fmt.Sprintf("%s", s.NameTW)
}

func (s *Station) setDistanceAwayFrom(latitude, longitude float64) {
	radius := 6378.137
	rad := math.Pi / 180.0
	lat1 := s.Latitude * rad
	lng1 := s.Longitude * rad
	lat2 := latitude * rad
	lng2 := longitude * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	s.Distance = dist * radius

	// log.Printf("DEBUG %s %g", s.NameTW, s.Distance)
}

// IData ...
type IData interface {
	GetIP() net.IP
	SetStation(Station)
}

// Map ...
func Map(data IData) {
	ip := data.GetIP()
	station := ipToTaipeiMRTStation(ip)
	data.SetStation(station)
}

func ipStrToNetIP(ip string) net.IP {
	netIP := make(net.IP, 0)
	for _, digit := range strings.Split(ip, ".") {
		b, _ := strconv.Atoi(digit)
		if b < 0 || b > 256 {
			panic("invalid ip")
		}
		netIP = append(netIP, byte(b))
	}
	if len(netIP) != 4 {
		panic("invalid ip")
	}
	return netIP
}

type payloadKeyCDN struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		Geo struct {
			Host          string  `json:"host"`
			IP            string  `json:"ip"`
			RDNS          string  `json:"rdns"`
			ASN           string  `json:"asn"`
			ISP           string  `json:"isp"`
			CountryName   string  `json:"country_name"`
			CountryCode   string  `json:"country_code"`
			RegionName    string  `json:"region_name"`
			RegionCode    string  `json:"region_code"`
			City          string  `json:"city"`
			PostalCode    string  `json:"postal_code"`
			ContinentName string  `json:"continent_name"`
			ContinentCode string  `json:"continent_code"`
			Latitude      float64 `json:"latitude"`
			Longitude     float64 `json:"longitude"`
			MetroCode     string  `json:"metro_code"`
			Timezone      string  `json:"timezone"`
			Datetime      string  `json:"datetime"`
		} `json:"geo"`
	} `json:"data"`
}

func ipToLatitudeLongitude(ip net.IP) (float64, float64) {
	url := fmt.Sprintf("https://tools.keycdn.com/geo.json?host=%s", ip)
	// log.Printf("DEBUG %s", url)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("wrong http status code")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	payload := &payloadKeyCDN{}
	err = json.Unmarshal(body, payload)
	if err != nil {
		panic(err)
	}

	return payload.Data.Geo.Latitude, payload.Data.Geo.Longitude

	// API CALL EXAMPLE
	//
	// curl -X GET https://tools.keycdn.com/geo.json?host=${IP OR HOSTNAME}
	//
	//	{
	//		"status": "success",
	//		"description": "Data successfully received.",
	//		"data": {
	//			"geo": {
	//				"host": "101.15.22.238",
	//				"ip": "101.15.22.238",
	//				"rdns": "101.15.22.238",
	//				"asn": 24158,"isp":
	//				"Taiwan Mobile Co., Ltd.",
	//				"country_name": "Taiwan",
	//				"country_code": "TW",
	//				"region_name": "Taipei City",
	//				"region_code": "TPE",
	//				"city": "Taipei",
	//				"postal_code": null,
	//				"continent_name": "Asia",
	//				"continent_code": "AS",
	//				"latitude": 25.0478,
	//				"longitude": 121.5318,
	//				"metro_code": null,
	//				"timezone":	"Asia\/Taipei",
	//				"datetime": "2018-10-08 09:50:30"
	//			}
	//		}
	//	}
}

type payloadGoogle struct {
	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
	} `json:"location"`
	Accuracy float64 `json:"accuracy"`
}

func googleMyLatitudeLongitude() (float64, float64) {
	url := fmt.Sprintf("https://www.googleapis.com/geolocation/v1/geolocate?key=%s", os.Getenv("GoogleGeolocationAPIKey"))
	// log.Printf("DEBUG %s", url)

	res, err := http.Post(url, "Content-Type: application/json", strings.NewReader("{}"))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("wrong http status code")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// log.Printf("DEBUG %s", body)

	payload := &payloadGoogle{}
	err = json.Unmarshal(body, payload)
	if err != nil {
		panic(err)
	}

	// log.Printf("DEBUG %g %g", payload.Location.Latitude, payload.Location.Longitude)
	return payload.Location.Latitude, payload.Location.Longitude

	// API CALL EXAMPLE
	//
	// curl -X POST https://www.googleapis.com/geolocation/v1/geolocate?key=${GoogleGeolocationAPIKey} -H Content-Type: application/json -d "{}"
	//
	//	{
	//		"location": {
	//			"lat": 25.0626048,
	//			"lng": 121.6569344
	//		},
	//		"accuracy": 464.0
	//	}
}

func ipToTaipeiMRTStation(ip net.IP) (nearStation Station) {
	latitude, longitude := ipToLatitudeLongitude(ip)
	nearStation = findNearTaipeiMRTStation(latitude, longitude)
	return
}

func googleMyTaipeiMRTStation() (nearStation Station) {
	latitude, longitude := googleMyLatitudeLongitude()
	nearStation = findNearTaipeiMRTStation(latitude, longitude)
	return
}

func findNearTaipeiMRTStation(latitude, longitude float64) (nearStation Station) {
	stations := []*Station{
		&Station{NameTW: "南港展覽館", Latitude: 25.0553846, Longitude: 121.6182655},
		&Station{NameTW: "昆陽", Latitude: 25.0501585, Longitude: 121.593423},
		&Station{NameTW: "後山埤", Latitude: 25.045054, Longitude: 121.582522},
		&Station{NameTW: "永春", Latitude: 25.0407066, Longitude: 121.5765839},
		&Station{NameTW: "市政府", Latitude: 25.041146, Longitude: 121.565963},
		&Station{NameTW: "國父紀念館", Latitude: 25.039012, Longitude: 121.557739},
		&Station{NameTW: "忠孝敦化", Latitude: 25.0416218, Longitude: 121.5516552},
		&Station{NameTW: "忠孝復興", Latitude: 25.0553846, Longitude: 25.0553846},
		&Station{NameTW: "忠孝新生", Latitude: 25.0424433, Longitude: 121.5332498},
		&Station{NameTW: "善導寺", Latitude: 25.0441055, Longitude: 121.524231},
		&Station{NameTW: "台北車站", Latitude: 25.0477505, Longitude: 121.5170599},
		&Station{NameTW: "西門", Latitude: 25.0421203, Longitude: 121.5076592},
		&Station{NameTW: "龍山寺", Latitude: 25.035279, Longitude: 121.499826},
		&Station{NameTW: "江子翠", Latitude: 25.033971, Longitude: 121.473281},
		&Station{NameTW: "新埔", Latitude: 25.0241811, Longitude: 121.4682288},
		&Station{NameTW: "板橋", Latitude: 25.0145999, Longitude: 121.4625299},
		&Station{NameTW: "府中", Latitude: 25.0083954, Longitude: 121.4589482},
		&Station{NameTW: "亞東醫院", Latitude: 24.998037, Longitude: 121.452514},
		&Station{NameTW: "海山", Latitude: 24.9854279, Longitude: 121.4486084},
		&Station{NameTW: "土城", Latitude: 24.9732244, Longitude: 121.4441573},
		&Station{NameTW: "永寧", Latitude: 24.9668026, Longitude: 121.4363174},
		&Station{NameTW: "頂埔", Latitude: 24.959332, Longitude: 121.4198624},
		&Station{NameTW: "淡水", Latitude: 25.167817, Longitude: 121.44556},
		&Station{NameTW: "紅樹林", Latitude: 25.154042, Longitude: 121.458871},
		&Station{NameTW: "竹圍", Latitude: 25.136558, Longitude: 121.46009},
		&Station{NameTW: "關渡", Latitude: 25.124923, Longitude: 121.465515},
		&Station{NameTW: "忠義", Latitude: 25.131058, Longitude: 121.473226},
		&Station{NameTW: "復興崗", Latitude: 25.138198, Longitude: 121.491851},
		&Station{NameTW: "北投", Latitude: 25.1313054, Longitude: 121.4991801},
		&Station{NameTW: "新北投", Latitude: 25.1369429, Longitude: 121.5026492},
		&Station{NameTW: "奇岩", Latitude: 25.126112, Longitude: 121.5010071},
		&Station{NameTW: "唭哩岸", Latitude: 25.120872, Longitude: 121.506252},
		&Station{NameTW: "石牌", Latitude: 25.114523, Longitude: 121.515559},
		&Station{NameTW: "明德", Latitude: 25.10972, Longitude: 121.518848},
		&Station{NameTW: "芝山", Latitude: 25.1030598, Longitude: 121.5225139},
		&Station{NameTW: "士林", Latitude: 25.093535, Longitude: 121.526229},
		&Station{NameTW: "劍潭", Latitude: 25.084873, Longitude: 121.525077},
		&Station{NameTW: "圓山", Latitude: 25.071353, Longitude: 121.520118},
		&Station{NameTW: "民權西路", Latitude: 25.0628786, Longitude: 121.519346},
		&Station{NameTW: "雙連", Latitude: 25.0579015, Longitude: 121.5206032},
		&Station{NameTW: "中山", Latitude: 25.0529046, Longitude: 121.5203406},
		&Station{NameTW: "台大醫院", Latitude: 25.041471, Longitude: 121.518577},
		&Station{NameTW: "中正紀念堂", Latitude: 25.032729, Longitude: 121.51827},
		&Station{NameTW: "東門", Latitude: 25.0340535, Longitude: 121.5289343},
		&Station{NameTW: "大安森林公園", Latitude: 25.0334788, Longitude: 121.535352},
		&Station{NameTW: "大安", Latitude: 25.0332255, Longitude: 121.5439082},
		&Station{NameTW: "信義安和", Latitude: 25.033116, Longitude: 121.552798},
		&Station{NameTW: "台北101/世貿", Latitude: 25.0330575, Longitude: 121.5631862},
		&Station{NameTW: "象山", Latitude: 25.032939, Longitude: 121.5688409},
		&Station{NameTW: "南港軟體園區", Latitude: 25.0599419, Longitude: 121.6159291},
		&Station{NameTW: "東湖", Latitude: 25.0671258, Longitude: 121.6117183},
		&Station{NameTW: "葫洲", Latitude: 25.0727158, Longitude: 121.6071548},
		&Station{NameTW: "大湖公園", Latitude: 25.083009, Longitude: 121.605262},
		&Station{NameTW: "內湖", Latitude: 25.083718, Longitude: 121.594418},
		&Station{NameTW: "文德", Latitude: 25.0785484, Longitude: 121.5855236},
		&Station{NameTW: "港墘", Latitude: 25.080038, Longitude: 121.575102},
		&Station{NameTW: "西湖", Latitude: 25.082133, Longitude: 121.567213},
		&Station{NameTW: "劍南路", Latitude: 25.084853, Longitude: 121.555592},
		&Station{NameTW: "大直", Latitude: 25.079476, Longitude: 121.546894},
		&Station{NameTW: "松山機場", Latitude: 25.063021, Longitude: 121.5517715},
		&Station{NameTW: "中山國中", Latitude: 25.060848, Longitude: 121.544226},
		&Station{NameTW: "南京復興", Latitude: 25.0520954, Longitude: 121.5443516},
		&Station{NameTW: "科技大樓", Latitude: 25.026124, Longitude: 121.543437},
		&Station{NameTW: "六張犁", Latitude: 25.023698, Longitude: 121.55162},
		&Station{NameTW: "麟光", Latitude: 25.015464, Longitude: 121.561065},
		&Station{NameTW: "辛亥", Latitude: 25.005475, Longitude: 121.557107},
		&Station{NameTW: "萬芳醫院", Latitude: 24.9993886, Longitude: 121.5582008},
		&Station{NameTW: "萬芳社區", Latitude: 25.00225, Longitude: 121.569992},
		&Station{NameTW: "木柵", Latitude: 24.99824, Longitude: 121.573145},
		&Station{NameTW: "動物園", Latitude: 24.998259, Longitude: 121.57943},
		&Station{NameTW: "松山", Latitude: 25.0498846, Longitude: 121.5775295},
		&Station{NameTW: "南京三民", Latitude: 25.051462, Longitude: 121.560447},
		&Station{NameTW: "台北小巨蛋", Latitude: 25.0514463, Longitude: 121.5507862},
		&Station{NameTW: "松江南京", Latitude: 25.0512257, Longitude: 121.5327387},
		&Station{NameTW: "北門", Latitude: 25.0490258, Longitude: 121.5099912},
		&Station{NameTW: "小南門", Latitude: 25.035534, Longitude: 121.51088},
		&Station{NameTW: "中正紀念堂", Latitude: 25.032729, Longitude: 121.51827},
		&Station{NameTW: "古亭", Latitude: 25.0261342, Longitude: 121.522922},
		&Station{NameTW: "台電大樓", Latitude: 25.017557, Longitude: 121.531394},
		&Station{NameTW: "公館", Latitude: 25.011943, Longitude: 121.537818},
		&Station{NameTW: "萬隆", Latitude: 24.999282, Longitude: 121.540237},
		&Station{NameTW: "景美", Latitude: 24.993524, Longitude: 121.540817},
		&Station{NameTW: "大坪林", Latitude: 24.9827862, Longitude: 121.5410329},
		&Station{NameTW: "七張", Latitude: 24.9749794, Longitude: 121.5431015},
		&Station{NameTW: "新店區公所", Latitude: 24.9673862, Longitude: 121.5410004},
		&Station{NameTW: "新店", Latitude: 24.95894, Longitude: 121.538581},
		&Station{NameTW: "南勢角", Latitude: 24.9899178, Longitude: 121.5088501},
		&Station{NameTW: "景安", Latitude: 24.9937115, Longitude: 121.5046146},
		&Station{NameTW: "永安市場", Latitude: 25.003881, Longitude: 121.51332},
		&Station{NameTW: "頂溪", Latitude: 25.014218, Longitude: 121.5153677},
		&Station{NameTW: "行天宮", Latitude: 25.059717, Longitude: 121.533184},
		&Station{NameTW: "中山國小", Latitude: 25.062693, Longitude: 121.526418},
		&Station{NameTW: "大橋頭", Latitude: 25.0632553, Longitude: 121.5127182},
		&Station{NameTW: "三重國小", Latitude: 25.06945, Longitude: 121.498291},
		&Station{NameTW: "三和國中", Latitude: 25.07542, Longitude: 121.488685},
		&Station{NameTW: "徐匯中學", Latitude: 25.079547, Longitude: 121.481262},
		&Station{NameTW: "三民高中", Latitude: 25.085457, Longitude: 121.473389},
		&Station{NameTW: "蘆洲", Latitude: 25.091553, Longitude: 121.464471},
		&Station{NameTW: "台北橋", Latitude: 25.063059, Longitude: 121.499916},
		&Station{NameTW: "菜寮", Latitude: 25.060274, Longitude: 121.492156},
		&Station{NameTW: "三重", Latitude: 25.055791, Longitude: 121.484725},
		&Station{NameTW: "先嗇宮", Latitude: 25.046493, Longitude: 121.471916},
		&Station{NameTW: "頭前庄", Latitude: 25.039705, Longitude: 121.461746},
		&Station{NameTW: "新莊", Latitude: 25.036125, Longitude: 121.452468},
		&Station{NameTW: "輔大", Latitude: 25.0328241, Longitude: 121.4361096},
		&Station{NameTW: "丹鳳", Latitude: 25.0290509, Longitude: 121.422421},
		&Station{NameTW: "迴龍", Latitude: 25.02254, Longitude: 121.412734},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(stations))
	for _, station := range stations {
		go func(station *Station) {
			defer wg.Done()
			station.setDistanceAwayFrom(latitude, longitude)
		}(station)
	}
	wg.Wait()

	minDistance := math.MaxFloat64
	for _, station := range stations {
		if minDistance > station.Distance {
			minDistance = station.Distance
			nearStation = *station
		}
	}

	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	if len(os.Args) == 1 {
		fmt.Println(googleMyTaipeiMRTStation())
	} else {
		ip := os.Args[1]
		netIP := ipStrToNetIP(ip)
		fmt.Println(ipToTaipeiMRTStation(netIP))
	}
}
