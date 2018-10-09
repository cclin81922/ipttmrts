package ipttmrts

import (
	"fmt"
)

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
