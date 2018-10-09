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

package ipttmrts

import (
	"fmt"
	"net"
)

type example struct {
	ip      net.IP
	station string
}

func (e example) GetIP() net.IP {
	return e.ip
}

func (e *example) SetStation(s Station) {
	e.station = s.NameTW
}

func (e example) String() string {
	return fmt.Sprintf("ip %s | station %s", e.ip, e.station)
}

func ExampleMap() {
	data := &example{ip: []byte{125, 227, 32, 90}}
	Map(data)
	fmt.Println(data)
	//Output:
	//ip 125.227.32.90 | station 松江南京
}

func ExampleFindNearTaipeiMRTStation() {
	fmt.Println(FindNearTaipeiMRTStation(25.054361, 121.619107))
	//Output:
	//南港展覽館
}

func ExampleIPToTaipeiMRTStation() {
	fmt.Println(IPToTaipeiMRTStation([]byte{125, 227, 32, 90}))
	//Output:
	//松江南京
}
