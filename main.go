package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// Station ...
type Station struct {
	NameTW    string
	NameEN    string
	Latitude  float64
	Longitude float64
}

func (s Station) String() string {
	return fmt.Sprintf("%s", s.NameTW)
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

func merge(distance ...<-chan float64) <-chan float64 {
	out := make(chan float64)

	// TODO

	return out
}

func distanceOfTwoPointsByLatitudeLongitude(latitude1, longitude1, latitude2, longitude2 float64) <-chan float64 {
	out := make(chan float64)

	// TODO

	return out
}

func ipToLatitudeLongitude(ip net.IP) (float64, float64) {
	// TODO
	return 0, 0
}

func ipToTaipeiMRTStation(ip net.IP) Station {
	stations := []Station{
		{"南港展覽館", "", 0, 0},
		{"昆陽", "", 0, 0},
		{"後山埤", "", 0, 0},
		{"永春", "", 0, 0},
		{"市政府", "", 0, 0},
		{"國父紀念館", "", 0, 0},
		{"忠孝敦化", "", 0, 0},
		{"忠孝復興", "", 0, 0},
		{"忠孝新生", "", 0, 0},
		{"善導寺", "", 0, 0},
		{"台北車站", "", 0, 0},
		{"西門", "", 0, 0},
		{"龍山寺", "", 0, 0},
		{"江子翠", "", 0, 0},
		{"新埔", "", 0, 0},
		{"板橋", "", 0, 0},
		{"府中", "", 0, 0},
		{"亞東醫院", "", 0, 0},
		{"海山", "", 0, 0},
		{"土城", "", 0, 0},
		{"永寧", "", 0, 0},
		{"頂埔", "", 0, 0},
		{"淡水", "", 0, 0},
		{"紅樹林", "", 0, 0},
		{"竹圍", "", 0, 0},
		{"關渡", "", 0, 0},
		{"忠義", "", 0, 0},
		{"復興崗", "", 0, 0},
		{"北投", "", 0, 0},
		{"新北投", "", 0, 0},
		{"奇岩", "", 0, 0},
		{"唭哩岸", "", 0, 0},
		{"石牌", "", 0, 0},
		{"明德", "", 0, 0},
		{"芝山", "", 0, 0},
		{"士林", "", 0, 0},
		{"劍潭", "", 0, 0},
		{"圓山", "", 0, 0},
		{"民權西路", "", 0, 0},
		{"雙連", "", 0, 0},
		{"中山", "", 0, 0},
		{"台大醫院", "", 0, 0},
		{"中正紀念堂", "", 0, 0},
		{"東門", "", 0, 0},
		{"大安森林公園", "", 0, 0},
		{"大安", "", 0, 0},
		{"信義安和", "", 0, 0},
		{"台北101/世貿", "", 0, 0},
		{"象山", "", 0, 0},
	}
	_ = stations

	latitude, longitude := ipToLatitudeLongitude(ip)

	distances := make([]<-chan float64, 0)
	for _, station := range stations {
		distance := distanceOfTwoPointsByLatitudeLongitude(station.Latitude, station.Longitude, latitude, longitude)
		distances = append(distances, distance)
	}

	var minDistance float64
	for distance := range merge(distances...) {
		if minDistance > distance {
			minDistance = distance
		}
	}

	// TODO return the station of minimum distance
	return Station{NameTW: "西門", NameEN: "Ximen"}
}

func main() {
	ip := os.Args[1] // TODO error handling
	netIP := ipStrToNetIP(ip)
	fmt.Println(ipToTaipeiMRTStation(netIP))
}
