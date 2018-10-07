# Command Line Usage

```
go get github.com/cclin81922/ipttmrts

ipttmrts 101.15.22.238
``` 

# Package Usage

```
import "github.com/cclin81922/ipttmrts"

func demo(data ipttmrts.IData) {
    ipttmrts.Map(data)
}
```

# About IData interface and Map function

```
type IData interface {
    GetIP() net.IP
    SetStation(Station)
}

func Map(data IData) {
    ip := data.GetIP()
    station  := ipToTaipeiMRTStation(ip)
    data.SetStation(station)
}
```
