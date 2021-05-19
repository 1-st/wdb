package util

import (
	"time"
)

func DurationHour(a, b time.Time) float64 {
	return float64(b.Unix()-a.Unix()) / 3600
}

func HourToNow(fmt string) float64 {
	before, _ := time.Parse(time.RFC3339, fmt)
	return DurationHour(before, time.Now())
}

