package time

import (
	"context"
	"time"
)

//DayBeginStamp ...
func DayBeginStamp(now int64) int64 {

	_, offset := time.Now().Zone()
	return now - (now+int64(offset))%int64(3600*24)

}

//HourBeginStamp ...
func HourBeginStamp(now int64) int64 {

	_, offset := time.Now().Zone()
	return now - (now+int64(offset))%int64(3600)

}

// DayBeginStampFromStr 获取指定天的时间范围
// 天格式 2006-01-02
// 为空时候返回当天的
func DayBeginStampFromStr(day string) (int64, error) {
	nowt := time.Now()
	now := nowt.Unix()

	var begin int64
	if len(day) > 0 {
		tm, err := time.ParseInLocation("2006-01-02", day, nowt.Location())
		if err != nil {
			return 0, err
		}

		begin = tm.Unix()

	} else {
		begin = DayBeginStamp(now)

	}

	return begin, nil

}

//WeekScope ..
// 指定时间所在周的范围
func WeekScope(stamp int64) (int64, int64) {
	now := time.Unix(stamp, 0)
	weekday := time.Duration(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := now.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	begin := currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour).Unix()
	return begin, begin + 24*3600*7 - 1
}

//MonthScope ...
// 指定时间所在月的范围
func MonthScope(stamp int64) (int64, int64) {
	now := time.Unix(stamp, 0)
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth.Unix(), lastOfMonth.Unix() + 3600*24 - 1
}

//RunTimeStat ...
type RunTimeStat struct {
	//logkey string
	since time.Time
}

//Millisecond ...
func (m *RunTimeStat) Millisecond() int64 {
	return m.Microsecond() / 1000
}

//Microsecond ...
func (m *RunTimeStat) Microsecond() int64 {
	return m.Duration().Nanoseconds() / 1000

}

//Nanosecond ...
func (m *RunTimeStat) Nanosecond() int64 {
	return m.Duration().Nanoseconds()
}

//Duration ...
func (m *RunTimeStat) Duration() time.Duration {
	return time.Since(m.since)
}

//Reset ...
func (m *RunTimeStat) Reset() {
	m.since = time.Now()
}

//NewTimeStat ...
func NewTimeStat() *RunTimeStat {
	return &RunTimeStat{
		//logkey: key,
		since: time.Now(),
	}
}

type Duration time.Duration

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

// Shrink will decrease the duration by comparing with context's timeout duration and return new timeout\context\CancelFunc.
func (d Duration) Shrink(c context.Context) (Duration, context.Context, context.CancelFunc) {
	if deadline, ok := c.Deadline(); ok {
		if ctimeout := time.Until(deadline); ctimeout < time.Duration(d) {
			// deliver small timeout
			return Duration(ctimeout), c, func() {}
		}
	}
	ctx, cancel := context.WithTimeout(c, time.Duration(d))
	return d, ctx, cancel
}
