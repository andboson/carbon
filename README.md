## Carbon

Is a time toolkit for golang, go implementation of [http://carbon.nesbot.com/](http://carbon.nesbot.com/) (PHP time toolkit)


## Install

```
go get -u github.com/andboson/carbon
```

### Usage

```go
import "github.com/andboson/carbon"

time.Now() // 2013-11-18 17:51:49

c := carbon.Now()   //  2016-02-10 11:24:50.761043525 +0200 EET
c.AddDays(2) // 2016-02-12 11:24:50.761043525 +0200 EET
c.SubDay() // 2016-02-11 11:25:38.766784153 +0200 EET
c.ToDateTimeString() //2016-02-11 11:26:18
c.SetTZ("UTC") // 2016-02-11 09:27:37.746725832 +0000 UTC
time2 := *c
time2.AddDays(1)
c.DiffInHours(time2) // -24
time3 := *c
time3.SubDay()
c.Between(time3, time2) // true
stamp := "2015-01-25 15:04:55"
time4, _ := CreateFrom(stamp) // 2015-01-25 15:04:55 +0000 UTC
c.Gt(*time4) // true

```

see [https://github.com/andboson/carbon/blob/master/carbon_test.go](https://github.com/andboson/carbon/blob/master/carbon_test.go) for more samples
