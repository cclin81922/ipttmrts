package main

import (
	"fmt"
	"net"
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

	// TODO
	s.Distance = 0

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
		b, _ := strconv.Atoi(digit) // TODO error handling
		netIP = append(netIP, byte(b))
	}
	return netIP
}

func ipToLatitudeLongitude(ip net.IP) (float64, float64) {
	// TODO
	return 25.0478, 121.5320
}

func ipToTaipeiMRTStation(ip net.IP) Station {
	stations := []*Station{
		&Station{NameTW: "南港展覽館", Latitude: 25.0553846, Longitude: 25.0553846},
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
	}

	latitude, longitude := ipToLatitudeLongitude(ip)

	wg := sync.WaitGroup{}
	wg.Add(len(stations))
	for _, station := range stations {
		go func(station *Station) {
			defer wg.Done()
			station.setDistanceAwayFrom(latitude, longitude)
		}(station)
	}
	wg.Wait()

	var minDistance float64
	var nearStation *Station
	for _, station := range stations {
		if minDistance > station.Distance {
			minDistance = station.Distance
			nearStation = station
		}
	}

	return *nearStation
}

func main() {
	ip := os.Args[1] // TODO error handling
	netIP := ipStrToNetIP(ip)
	fmt.Println(ipToTaipeiMRTStation(netIP))
}
