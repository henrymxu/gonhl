package gonhl

import (
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

func CreateDateFromString(dateString string) (time.Time, error) {
	return time.Parse(dateLayout, dateString)
}

func CreateDateFromTime(date time.Time) string {
	return fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
}

// Used for the expand query in various endpoints
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

func combineStringArray(array []string) string {
	return strings.Join(array, ",")
}

func combineIntArray(array []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}

func createSeasonString(season int) string {
	return fmt.Sprintf("%d%d", season, season + 1)
}

func createTimeStamp(date time.Time) string {
	return fmt.Sprintf("%d%d%d_%d%d%d", date.Year(),
		date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())
}