## Carbon

Is a time toolkit for golang, go implementation of [http://carbon.nesbot.com/](http://carbon.nesbot.com/) (PHP time toolkit)

## Changelog

- Added json Unmarshal function. Now you can use carbon.Carbon instead time.Time in structures. It can parse many time formats
- added create from string datetime functon
- added create from time.Time function
- init

## Install

```
go get -u github.com/andboson/carbon
```

### Usage

```go
import "github.com/andboson/carbon"

time.Now() // 2013-11-18 17:51:49

	log.Printf("%s", Now()) // 2016-02-10 13:22:13.566251776 +0200 EET
	log.Printf("%s", Now().AddDays(2)) // 2016-02-12 13:22:13.56669336 +0200 EET
	log.Printf("%s", Now().SubDay())  // 2016-02-09 13:22:13.566810314 +0200 EETT
	log.Printf("%s", Now().ToDateTimeString()) // 2016-02-10 13:22:13
	log.Printf("%s", Now().SetTZ("UTC")) // 2016-02-10 11:22:13.567020044 +0000 UTC
	time2 := Now().AddDay()
	log.Printf("%s", Now().DiffInHours(*time2)) // -24
	time3 := Now().SubDay()
	log.Printf("%s", Now().Between(*time3, *time2)) // true
	stamp := "2015-01-25 15:04:55"
	time4, _ := CreateFrom(stamp)  // 2015-01-25 15:04:55 +0000 UTC
	log.Printf("%s", Now().Gt(*time4)) // true
	log.Printf("%s", Now().StartOfHour()) // 2016-02-10 14:00:00 +0200 EET
	log.Printf("%s", Now().EndOfHour()) // 2016-02-10 14:59:59 +0200 EET
	log.Printf("%s", Now().StartOfDay()) // 2016-02-10 00:00:00 +0200 EET
	log.Printf("%s", Now().EndOfDay()) // 2016-02-10 23:59:59 +0200 EET
	log.Printf("%s", Now().StartOfWeek()) // 2016-02-08 00:00:00 +0200 EET
	log.Printf("%s", Now().EndOfWeek()) // 2016-02-14 23:59:59 +0200 EET
	log.Printf("%s", Now().StartOfMonth()) // 2016-02-01 00:00:00 +0200 EET
	log.Printf("%s", Now().EndOfMonth()) // 2016-02-29 23:59:59 +0200 EET
	log.Printf("%s", Now().StartOfYear()) // 2016-01-01 00:00:00 +0200 EET
	log.Printf("%s", Now().EndOfYear()) // 2016-12-31 23:59:59 +0200 EET

```

see [https://github.com/andboson/carbon/blob/master/carbon_test.go](https://github.com/andboson/carbon/blob/master/carbon_test.go) for samples
