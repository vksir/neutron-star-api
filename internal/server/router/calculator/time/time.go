package time

import (
	"time"
)

type Time struct {
	Year int `json:"year"`
}

func Add(year, month, day, hour, min, sec int, d time.Duration) time.Time {
	t := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC)
	return t.Add(d)
}

func Sub() {

}
