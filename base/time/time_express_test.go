// 在处理时间时始终使用 "time" 包，因为它有助于以更安全、更准确的方式处理时间
package time

import (
	"fmt"
	"testing"
	"time"
)

/*
- 对外部系统使用 time.Time 和 time.Duration
尽可能在与外部系统的交互中使用 time.Duration 和 time.Time 例如 :
Command-line 标志: flag 通过 time.ParseDuration 支持 time.Duration
JSON: encoding/json 通过其 UnmarshalJSON method 方法支持将 time.Time 编码为 RFC 3339 字符串
SQL: database/sql 支持将 DATETIME 或 TIMESTAMP 列转换为 time.Time，如果底层驱动程序支持则返回
YAML: gopkg.in/yaml.v2 支持将 time.Time 作为 RFC 3339 字符串，并通过 time.ParseDuration 支持 time.Duration。
*/

// 当不能在这些交互中使用 time.Duration 时，请使用 int 或 float64，并在字段名称中包含单位。 eg: {"intervalMillis": 2000}
type Config struct {
	IntervalMillis int `json:"intervalMillis"`
}

// 当不能在这些交互中使用 time.Time 时，除非达成一致，否则使用 string 和 RFC 3339 中定义的格式时间戳。默认情况下，Time.UnmarshalText 使用此格式，并可通过 time.RFC3339 在 Time.Format 和 time.Parse 中使用。
func TestGetTimestamp(t *testing.T) {

	// createTimeEnd := time.Unix(0, 0).Format("2006-01-02 15:04:05")
	// fmt.Println("--", createTimeEnd)

	noe := time.Now()
	// 获取24h后的时刻
	// maybeNewDay := noe.Add(24 * time.Hour)
	// fmt.Println(maybeNewDay)

	// 获取上一个日历日： 注意内部逻辑问题：如：8月31号 加一个月后 会变成 9月31号(不存在) => 10月1号
	newDay := noe.AddDate(0 /* years */, 1 /* months */, 0 /* days */)
	fmt.Println("==", newDay)

	nextDay := time.Date(noe.Year(), noe.Month(), 1, 0, 0, 0, 0, noe.Location())
	nextMonthFirstDay := nextDay.AddDate(0 /* years */, 1 /* months */, 0 /* days */)
	fmt.Println("==", nextMonthFirstDay)

	//
	// //获取下一个日历日10am
	//now := time.Now()
	// next := now.Add(time.Hour * 24)
	// next = time.Date(next.Year(), next.Month(), next.Day(), 10, 0, 0, 0, next.Location())
	// fmt.Println("==" , newDay)
	//
	//
	// duration := time.Now().Sub(time.Unix(time.Now().Unix(), 0))
	// if duration > 30*time.Minute {
	//
	// }

}

/*
日期字符串、Unix => Time
*/
func TestCovert(t *testing.T) {
	// date -> timestamp
	dateStr := "2020-08-11"
	begindate, err := time.ParseInLocation("2006-01-02", dateStr, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(" %d \n--\n", begindate.Unix())

	// timestamp -> date
	timestamp := int64(1595284492) // 2020-07-21 06:34:52
	tt := time.Unix(timestamp, 0)
	fmt.Println(tt.Format("2006-01-02 15:04:05"))
}
