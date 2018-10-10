# Installation

```
go get -u github.com/cclin81922/ipttmrts/cmd/ipttmrts
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
ipttmrts -ip=101.15.22.238

// output the nearest MRT station
```

To output debug message as well, use flag `-logtostderr=true -v=2`

# Package Usage

```
import "github.com/cclin81922/ipttmrts/pkg/ipttmrts"

func demo(data ipttmrts.IData) {
    ipttmrts.Map(data)
}
```

# For Developer

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

View API doc by web browser (offline)

```
godoc -http=:6060
open http://localhost:6060/pkg/github.com/cclin81922/ipttmrts/pkg/ipttmrts
```

View API doc by web browser (online)

```
open https://godoc.org/github.com/cclin81922/ipttmrts/pkg/ipttmrts
```

# Related Resources

* [The greate circle distance](http://wywu.pixnet.net/blog/post/26533759-%E7%B6%93%E7%B7%AF%E5%BA%A6%E8%A8%88%E7%AE%97%E8%B7%9D%E9%9B%A2%E5%85%AC%E5%BC%8F)
* [Online tool to calculate distance between two location by geographic coordinate](http://www.storyday.com/wp-content/uploads/2008/09/latlung_dis.html)
* [Open data of Taipei MRT stations locations with geographic coordinate](https://fusiontables.google.com/DataSource?docid=1QL2wqpruEvkPKhfb14Md9JMBzQIcKFFJ8wfAmORu#card:id=2)
* [Golang code gist which calculates distance between two location by geographic coordinate](https://blog.csdn.net/u013421629/article/details/72722714)
* [Golang code gist which demonstrates glog usage](https://gist.github.com/heatxsink/7221ebe499b0767d4784)
* [Google Geolocation API to get location geographic coordinate](https://developers.google.com/maps/documentation/geolocation/intro)
* [KeyCDN API to get locaiton geographic coordinate](https://tools.keycdn.com/geo)