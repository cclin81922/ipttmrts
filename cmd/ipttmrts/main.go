// Command ipttmrts ...
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cclin81922/ipttmrts/pkg/ipttmrts"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	if len(os.Args) == 1 {
		fmt.Println(ipttmrts.GoogleMyTaipeiMRTStation())
	} else {
		ip := os.Args[1]
		netIP := ipttmrts.IPStrToNetIP(ip)
		fmt.Println(ipttmrts.IPToTaipeiMRTStation(netIP))
	}
}
