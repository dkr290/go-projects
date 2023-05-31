package datehelpers

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05"
)

func GetTimeNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {

	return GetTimeNow().Format(apiDateLayout)

}
