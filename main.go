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
	return 0, 0
}

func ipToTaipeiMRTStation(ip net.IP) Station {
	stations := []*Station{
		&Station{NameTW: "南港展覽館", Latitude: 0, Longitude: 0},
		&Station{NameTW: "昆陽", Latitude: 0, Longitude: 0},
		&Station{NameTW: "後山埤", Latitude: 0, Longitude: 0},
		&Station{NameTW: "永春", Latitude: 0, Longitude: 0},
		&Station{NameTW: "市政府", Latitude: 0, Longitude: 0},
		&Station{NameTW: "國父紀念館", Latitude: 0, Longitude: 0},
		&Station{NameTW: "忠孝敦化", Latitude: 0, Longitude: 0},
		&Station{NameTW: "忠孝復興", Latitude: 0, Longitude: 0},
		&Station{NameTW: "忠孝新生", Latitude: 0, Longitude: 0},
		&Station{NameTW: "善導寺", Latitude: 0, Longitude: 0},
		&Station{NameTW: "台北車站", Latitude: 0, Longitude: 0},
		&Station{NameTW: "西門", Latitude: 0, Longitude: 0},
		&Station{NameTW: "龍山寺", Latitude: 0, Longitude: 0},
		&Station{NameTW: "江子翠", Latitude: 0, Longitude: 0},
		&Station{NameTW: "新埔", Latitude: 0, Longitude: 0},
		&Station{NameTW: "板橋", Latitude: 0, Longitude: 0},
		&Station{NameTW: "府中", Latitude: 0, Longitude: 0},
		&Station{NameTW: "亞東醫院", Latitude: 0, Longitude: 0},
		&Station{NameTW: "海山", Latitude: 0, Longitude: 0},
		&Station{NameTW: "土城", Latitude: 0, Longitude: 0},
		&Station{NameTW: "永寧", Latitude: 0, Longitude: 0},
		&Station{NameTW: "頂埔", Latitude: 0, Longitude: 0},
		&Station{NameTW: "淡水", Latitude: 0, Longitude: 0},
		&Station{NameTW: "紅樹林", Latitude: 0, Longitude: 0},
		&Station{NameTW: "竹圍", Latitude: 0, Longitude: 0},
		&Station{NameTW: "關渡", Latitude: 0, Longitude: 0},
		&Station{NameTW: "忠義", Latitude: 0, Longitude: 0},
		&Station{NameTW: "復興崗", Latitude: 0, Longitude: 0},
		&Station{NameTW: "北投", Latitude: 0, Longitude: 0},
		&Station{NameTW: "新北投", Latitude: 0, Longitude: 0},
		&Station{NameTW: "奇岩", Latitude: 0, Longitude: 0},
		&Station{NameTW: "唭哩岸", Latitude: 0, Longitude: 0},
		&Station{NameTW: "石牌", Latitude: 0, Longitude: 0},
		&Station{NameTW: "明德", Latitude: 0, Longitude: 0},
		&Station{NameTW: "芝山", Latitude: 0, Longitude: 0},
		&Station{NameTW: "士林", Latitude: 0, Longitude: 0},
		&Station{NameTW: "劍潭", Latitude: 0, Longitude: 0},
		&Station{NameTW: "圓山", Latitude: 0, Longitude: 0},
		&Station{NameTW: "民權西路", Latitude: 0, Longitude: 0},
		&Station{NameTW: "雙連", Latitude: 0, Longitude: 0},
		&Station{NameTW: "中山", Latitude: 0, Longitude: 0},
		&Station{NameTW: "台大醫院", Latitude: 0, Longitude: 0},
		&Station{NameTW: "中正紀念堂", Latitude: 0, Longitude: 0},
		&Station{NameTW: "東門", Latitude: 0, Longitude: 0},
		&Station{NameTW: "大安森林公園", Latitude: 0, Longitude: 0},
		&Station{NameTW: "大安", Latitude: 0, Longitude: 0},
		&Station{NameTW: "信義安和", Latitude: 0, Longitude: 0},
		&Station{NameTW: "台北101/世貿", Latitude: 0, Longitude: 0},
		&Station{NameTW: "象山", Latitude: 0, Longitude: 0},
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
