//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

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
