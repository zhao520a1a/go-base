package time

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestTimeZone(t *testing.T) {
	//source := time.Now().Unix()
	//str := "2020-03-11 00:00:00"   //美国夏令时（3月11日至11月7日）
	str := "2020-01-01 00:00:00" //美国夏令时（3月11日至11月7日）
	fmt.Println("str: ", str)
	fmt.Println("-- Parse 把时间字符串转换为Time，默认时区是UTC时区 ")
	tm, _ := time.Parse(TIME_LAYOUT, str)
	source := tm.Unix()

	//设置环境变量，
	os.Setenv("ZONEINFO", "/Users/Golden/Documents/GoProjects/practice-demo/src/source/zoneinfo.zip")

	locationNames := []string{
		"Asia/Shanghai",   //	+08:00	+08:00
		"America/Creston", //		−07:00	−07:00

		//夏令时城市
		"America/Anchorage", //	−09:00	−08:00
		"America/Dawson",    //	−08:00	−07:00
	}

	for _, name := range locationNames {
		target, err := transformTimestamp(source, name)
		fmt.Printf("locationName:%s, source: %d, res: %d, err:%v", name, source, target, err)
	}
}
