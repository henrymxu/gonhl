package gonhl

import (
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

// CreateDateFromString converts a string representing a date with format `yyyy-mm-dd` to a time.Time object.
func CreateDateFromString(dateString string) (time.Time, error) {
	return time.Parse(dateLayout, dateString)
}

// CreateStringFromDate converts a time.Time object to a string representing a date with format `yyyy-mm-dd`.
func CreateStringFromDate(date time.Time) string {
	return fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
}

// expandQuery concatenates boolean flags to be used in HTTP queries.
func expandQuery(endpoint string, toExpand map[string]bool) string {
	expand := ""
	for key, value := range toExpand {
		if value {
			if len(expand) > 0 {
				expand += ","
			}
			expand += fmt.Sprintf("%s.%s", endpoint, key)
		}
	}
	return expand
}

// combineStringArray converts an array of strings into a comma separated string.
func combineStringArray(array []string) string {
	return strings.Join(array, ",")
}

// combineIntArray converts an array of ints into a comma seperated string.
func combineIntArray(array []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}

// createSeasonString converts a year into a string representing an NHL season.
func createSeasonString(season int) string {
	return fmt.Sprintf("%d%d", season, season + 1)
}

// createTimeStamp converts a time.Time object into a string representing a date with format `yyyymmdd_hhmmss`.
func createTimeStamp(date time.Time) string {
	return fmt.Sprintf("%d%d%d_%d%d%d", date.Year(),
		date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())
}