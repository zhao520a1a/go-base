package time

import (
	"fmt"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		//zone, offset := time.Now().In(location).Zone()
		lt, err := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		if err != nil {
			println(err.Error())
		}
		fmt.Println(locationName, lt)
		return lt, nil
	}
}

func transformTimestamp(timestamp int64, locationName string) (target int64, err error) {
	//设置时区
	location, err := time.LoadLocation(locationName)
	if err != nil {
		println(err.Error())
		return
	}

	//parse time
	tt := time.Unix(timestamp, 0).In(location)

	zone, offset := tt.Zone()
	tStr := tt.Format(TIME_LAYOUT)
	fmt.Printf("\n-- zone:%s, offset:%d tStr:%s \n ", zone, offset, tStr)

	target = tt.Unix()
	return
}
