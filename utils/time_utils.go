package utils

import (
	"time"
	"github.com/mergestat/timediff"
)

// Define the layout for parsing the date-time
const Layout = "2006-01-02 15:04:05.999999999 -0700 MST"

// ParseDate parses a date string into a time.Time object
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(Layout, dateStr)
}

// FormatTimeDiff calculates and formats the time difference from now
func FormatTimeDiff(createdTime time.Time) string {
	now := time.Now()
	duration := now.Sub(createdTime)
	return timediff.TimeDiff(now.Add(-duration))
}