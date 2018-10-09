# Installation

```
go get github.com/cclin81922/ipttmrts/cmd/ipttmrts
export PATH=$PATH:~/go/bin
```

# Command Line Usage

Usage 1

```
GoogleGeolocationAPIKey=... ipttmrts

// output the nearest MRT station
``` 

Usage 2

```
ipttmrts 101.15.22.238

// output the nearest MRT station
```

# Package Usage

```
import "github.com/cclin81922/ipttmrts/pkg/ipttmrts"

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

# For developer

Run all tests

```
go test github.com/cclin81922/ipttmrts/pkg/ipttmrts
```

Run selected tests e.g.,

```
go test github.com/cclin81922/ipttmrts/pkg/ipttmrts -run=TestFindNearTaipeiMRTStation/誠正國中
```

Run all benchmarks

```
go test github.com/cclin81922/ipttmrts/pkg/ipttmrts -bench=. -benchmem
```

Run selected benchmarks e.g.,

```
go test github.com/cclin81922/ipttmrts/pkg/ipttmrts -bench=BenchmarkFindNearTaipeiMRTStation/誠正國中 -benchmem
```

View API doc by terminal

```
go doc github.com/cclin81922/ipttmrts/pkg/ipttmrts
```

View API doc by web browser

```
godoc -http=:6060
open http://localhost:6060/pkg/github.com/cclin81922/ipttmrts/pkg/ipttmrts/
```

# Related Resources

* [Online tool to calculate distance between two location by geographic coordinate](http://www.storyday.com/wp-content/uploads/2008/09/latlung_dis.html)
* [Open data of Taipei MRT stations locations with geographic coordinate](https://fusiontables.google.com/DataSource?docid=1QL2wqpruEvkPKhfb14Md9JMBzQIcKFFJ8wfAmORu#card:id=2)
* [Goland code to calculate distance between two location by geographic coordinate](https://blog.csdn.net/u013421629/article/details/72722714)
* [Google Geolocation API to get location geographic coordinate](https://developers.google.com/maps/documentation/geolocation/intro)
* [KeyCDN API to get locaiton geographic coordinate](https://tools.keycdn.com/geo)