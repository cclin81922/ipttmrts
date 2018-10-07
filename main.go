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
	NameTW string
	NameEN string
}

func (s Station) String() string {
	return fmt.Sprintf("TW %s | EN %s", s.NameTW, s.NameEN)
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

func ipToTaipeiMRTStation(ip net.IP) Station {
	// TODO
	// ipToLatitudeLongitude
	// distanceOfTwoPointsByLatitudeLongitude
	// pipelineOfCalculateThenSort
	return Station{NameTW: "西門", NameEN: "Ximen"}
}

func main() {
	ip := os.Args[1] // TODO error handling
	netIP := ipStrToNetIP(ip)
	fmt.Println(ipToTaipeiMRTStation(netIP))
}
