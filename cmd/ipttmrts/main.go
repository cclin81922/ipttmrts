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

// Command ipttmrts outputs the nearest Taipei MRT station according to your current IP address.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cclin81922/ipttmrts/pkg/ipttmrts"
	"github.com/golang/glog"
)

var (
	flagIP string
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: ipttmrts -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.StringVar(&flagIP, "ip", "", "IP address which will be used for finding the nearest Taipei MRT station")
	flag.Usage = usage
	flag.Parse()
}

func main() {
	defer func() {
		glog.Flush()
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	if flagIP == "" {
		if station, err := ipttmrts.GoogleMyTaipeiMRTStation(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(station)
		}
		fmt.Println()
	} else {
		if station, err := ipttmrts.IPToTaipeiMRTStation(flagIP); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(station)
		}
	}

}
