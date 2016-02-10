package carbon

import (
	"log"
	"testing"
	"time"
)

var format = "2006-01-02 15:04:05.999999999"

func TestDiff(t *testing.T) {
	//hours
	time1 := Create(2001, 12, 01, 1, 0)
	time2 := Create(2001, 12, 01, 3, 0)
	diff := time1.DiffInHours(*time2)

	if diff != -2 {
		t.Errorf("Diff in hours", diff)
	}

	time1 = Create(2001, 12, 01, 1)
	time2 = Create(2001, 12, 01, 2)
	diff = time1.DiffInMinutes(*time2)

	if diff != -60 {
		t.Errorf("Diff in minutes", diff)
	}

	time1 = Create(2001, 12, 01, 1, 1)
	time2 = Create(2001, 12, 01, 1, 2)
	diff = time1.DiffInSeconds(*time2)

	if diff != -60 {
		t.Errorf("Diff in seconds", diff)
	}
}

func TestSub(t *testing.T) {
	//hours
	time1 := Now()
	time2 := Now().SubDay()
	diff := time1.DiffInHours(*time2)
	if diff != 24.00 {
		t.Errorf("SubDay", time2, diff)
	}
}

func TestFormat(t *testing.T) {
	time1 := Now()
	if time1.ToDateTimeString() != time1.Format(DATE_TIME_LAYOUT) {
		t.Errorf("Format", time1.ToDateTimeString())
	}
}

func TestChain(t *testing.T) {
	time1 := Create(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	time2 := *time1 //pass by value
	time2.SubDay()

	if time1.Equal(time2.Time) {
		t.Errorf("Error compare", time1, time2)
	}
}

func TestTZ(t *testing.T) {
	time1 := Now()
	time1.SetTZ(`UTC`)
	time2 := *time1 //pass by value
	time2.SetTZ(`Europe/Kiev`)

	if time1.Location().String() == time2.Location().String() {
		t.Errorf("Error change tz", time1, time2)
	}
}

func TestCreateFromString(t *testing.T) {
	stamp := "2015-01-25 15:04:55"
	time1, _ := CreateFrom(stamp)

	if time1.ToDateTimeString() != stamp {
		t.Errorf("Error create time", time1)
	}
}

func TestBetween(t *testing.T) {
	time1 := Create(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	time2 := *time1 //pass by value
	time3 := *time1 //pass by value
	time2.SubDay()
	time3.AddDay()

	if !time1.Gt(time2) {
		t.Errorf("Error compare time greater", time1, time2)
	}

	if !time1.Between(time2, time3) {
		t.Errorf("Error compare time between", time1, time2, time3)
	}
}

func Example() {
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
}
