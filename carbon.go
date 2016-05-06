package carbon

import (
	"math"
	"time"
	"strings"
)

const (
	DATE_TIME_LAYOUT = "2006-01-02 15:04:05"
)

// form @github.com/jinzhu/now and more
var TimeFormats = []string{
	DATE_TIME_LAYOUT,
	"1/2/2006", "1/2/2006 15:4:5", "2006-1-2 15:4:5", "2006-1-2 15:4", "2006-1-2", "1-2", "15:4:5", "15:4",
	"2013-02-03", "15:4:5 Jan 2, 2006 MST", "15:04:05", "2013/02/03",
	"2013-02-03 19:54:00 PST",
	time.RFC822, time.RubyDate, time.RFC822Z, time.RFC3339,
}

type Carbon struct {
	time.Time
}

// return Carbon with time Now
func Now() *Carbon {
	c := &Carbon{time.Now()}
	return c
}

// return Carbon with time Now
func CreateFromTime(timeObject time.Time) *Carbon {
	c := &Carbon{timeObject}
	return c
}

// create Carbon with year, month, day, hour, minute, second, nanosecond, int,  TZ string
// Create(2001, 12, 01, 15, 25, 55, 0, "UTC") = 2001-12-01 13:25:55 +0000 UTC
func Create(year, month, day int, args ...interface{}) *Carbon {
	var tz string
	Month := time.Month(month)
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	nanosecond := time.Now().Nanosecond()
	location := time.Now().Location()
	for key, val := range args {
		if key == 0 {
			hour = val.(int)
		}
		if key == 1 {
			minute = val.(int)
		}
		if key == 2 {
			second = val.(int)
		}
		if key == 3 {
			nanosecond = val.(int)
		}
		if key == 4 {
			tz = val.(string)
		}
	}
	c := &Carbon{time.Date(year, Month, day, hour, minute, second, nanosecond, location)}
	if tz != "" {
		c.SetTZ(tz)
	}
	return c
}

// create Carbon from time string, return in UTC, if tz not specified
// CreateFrom("2015-01-25 15:4:5") = 2015-01-25 15:04:55 +0000 UTC
func CreateFrom(stringDate string) (*Carbon, error) {
	var carbon = Now()
	var err error
	var tm time.Time
	for _, format := range TimeFormats {
		tm, err = time.Parse(format, stringDate)
		if err == nil {
			carbon.Time = tm
			break
		}
	}

	return carbon, err
}
// custom Unmarshal func form many times formats
func (t *Carbon) UnmarshalJSON(buf []byte) error {
	tt, err := CreateFrom(strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt.Time
	return nil
}

// set tx location, ex. "UTC"
func (c *Carbon) SetTZ(tz string) *Carbon {
	location, _ := time.LoadLocation(tz)
	if location != nil {
		c.Time = c.Time.In(location)
	}
	return c
}

// set time, ex. "23.59.29"
func (c *Carbon) SetTime(hours, minutes, seconds int) *Carbon {
	c.Time = time.Date(
		c.Year(),
		c.Month(),
		c.Day(),
		hours,
		minutes,
		seconds,
		c.Nanosecond(),
		c.Location(),
	)

	return c
}

func (c *Carbon) SubDay() *Carbon {
	return c.SubDays(1)
}

func (c *Carbon) SubDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -days)
	return c
}

func (c *Carbon) SubMonth(days int) *Carbon {
	return c.SubMonths(1)
}

func (c *Carbon) SubMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, -months, 0)
	return c
}

func (c *Carbon) SubYear(days int) *Carbon {
	return c.SubYears(1)
}

func (c *Carbon) SubYears(years int) *Carbon {
	c.Time = c.Time.AddDate(-years, 0, 0)
	return c
}

func (c *Carbon) AddDay() *Carbon {
	return c.AddDays(1)
}

func (c *Carbon) AddDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, days)
	return c
}

func (c *Carbon) AddMonth(days int) *Carbon {
	return c.AddMonths(1)
}

func (c *Carbon) AddMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, months, 0)
	return c
}

func (c *Carbon) AddYear(days int) *Carbon {
	return c.AddYears(1)
}

func (c *Carbon) AddYears(years int) *Carbon {
	c.AddDate(years, 0, 0)
	return c
}

func (c *Carbon) DiffInSeconds(from Carbon) int {
	return round(c.Sub(from.Time).Seconds())
}

func (c *Carbon) DiffInMinutes(from Carbon) int {
	return round(c.Sub(from.Time).Minutes())
}

func (c *Carbon) DiffInHours(from Carbon) int {
	return round(c.Sub(from.Time).Hours())
}

// Determines if the instance is equal to another
func (c *Carbon) Eq(another Carbon) bool {
	return c.Equal(another.Time)
}

// Determines if the instance is greater (after) than another
func (c *Carbon) Gt(another *Carbon) bool {
	return c.After(another.Time)
}

// Determines if the instance is less (Before) than another
func (c *Carbon) Lt(another *Carbon) bool {
	return c.Before(another.Time)
}

// Determines if the instance is greater than before and less than after
func (c *Carbon) Between(before, after *Carbon) bool {
	return c.After(before.Time) && c.Before(after.Time)
}

func (c *Carbon) StartOfHour() *Carbon {
	c.Time = c.Truncate(time.Hour)
	return c
}

func (c *Carbon) EndOfHour() *Carbon {
	c.StartOfHour()
	c.Time = c.Add(time.Hour - time.Second)
	return c
}

func (c *Carbon) StartOfDay() *Carbon {
	c.Time = c.StartOfHour().Add(-time.Hour * time.Duration(c.Hour()))
	return c
}

func (c *Carbon) EndOfDay() *Carbon {
	c.Time = c.StartOfDay().Add(time.Hour*time.Duration(24) - time.Second)
	return c
}

func (c *Carbon) StartOfWeek(firstDayOfWeekIsMonday ...bool) *Carbon {
	firstDay := time.Monday
	corrFirstDay := 1
	if len(firstDayOfWeekIsMonday) > 0 {
		if !firstDayOfWeekIsMonday[0] {
			firstDay = time.Sunday
			corrFirstDay = 0
		}
	}
	c.StartOfDay()
	if c.Weekday() != firstDay {
		c.Time = c.Add(-time.Hour * 24 * time.Duration(-corrFirstDay+int(c.Weekday())))
	}
	return c
}

func (c *Carbon) EndOfWeek() *Carbon {
	c.Time = c.StartOfWeek().Add(time.Hour*time.Duration(24*7) - time.Second)
	return c
}

func (c *Carbon) StartOfMonth() *Carbon {
	year := c.Year()
	Month := c.Month()
	location := time.Now().Location()
	c = &Carbon{time.Date(year, Month, 1, 0, 0, 0, 0, location)}

	return c
}

func (c *Carbon) EndOfMonth() *Carbon {
	c.Time = c.StartOfMonth().AddDate(0, 1, 0).Add(-time.Second)
	return c
}

func (c *Carbon) StartOfYear() *Carbon {
	year := c.Year()
	location := time.Now().Location()
	c = &Carbon{time.Date(year, time.Month(1), 1, 0, 0, 0, 0, location)}

	return c
}

func (c *Carbon) EndOfYear() *Carbon {
	c.Time = c.StartOfYear().AddDate(1, 0, 0).Add(-time.Second)
	return c
}

//formatters

// return string with DateTime format "2006-01-25 15:04:05"
func (c *Carbon) ToDateTimeString() string {
	return c.Format(DATE_TIME_LAYOUT)
}

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}
