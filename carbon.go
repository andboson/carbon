package carbon

import (
	"math"
	"time"
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

// set tx location, ex. "UTC"
func (c *Carbon) SetTZ(tz string) *Carbon {
	location, _ := time.LoadLocation(tz)
	if location != nil {
		c.Time = c.Time.In(location)
	}
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

func (c *Carbon) DiffInSeconds(from Carbon) float64 {
	return math.Ceil(c.Sub(from.Time).Seconds())
}

func (c *Carbon) DiffInMinutes(from Carbon) float64 {
	return math.Ceil(c.Sub(from.Time).Minutes())
}

func (c *Carbon) DiffInHours(from Carbon) float64 {
	return math.Ceil(c.Sub(from.Time).Hours())
}

// Determines if the instance is equal to another
func (c *Carbon) Eq(another Carbon) bool {
	return c.Equal(another.Time)
}

// Determines if the instance is greater (after) than another
func (c *Carbon) Gt(another Carbon) bool {
	return c.After(another.Time)
}

// Determines if the instance is less (Before) than another
func (c *Carbon) Lt(another Carbon) bool {
	return c.Before(another.Time)
}

// Determines if the instance is greater than before and less than after
func (c *Carbon) Between(before, after Carbon) bool {
	return c.After(before.Time) && c.Before(after.Time)
}

// return string with DateTime format "2006-01-25 15:04:05"
func (c *Carbon) ToDateTimeString() string {
	return c.Format(DATE_TIME_LAYOUT)
}
