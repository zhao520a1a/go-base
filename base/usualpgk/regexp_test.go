package usualpkg

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	// searchStr := "John is 18 year old handsome boy 2020.01.01!"
	//
	// pat1 := "[0-9]+.[0-9]+"
	// if ok, _ := regexp.Match(pat1, []byte(searchStr)); ok {
	//	fmt.Println("匹配到了!")
	// }
	//
	// //将正则通过 Compile 方法返回一个 Regexp 对象
	// pat2 := "[0-9]+"
	// re, _ := regexp.Compile(pat2)
	// str1 := re.ReplaceAllString(searchStr, "###")
	// fmt.Println("替换匹配到部分：", str1)
	//
	// f := func(s string) string {
	//	v, _ := strconv.ParseFloat(s, 32)
	//	return strconv.FormatFloat(v*2, 'f', 2, 32)
	// }
	// str2 := re.ReplaceAllStringFunc(searchStr, f)
	// fmt.Println("参数为函数时:", str2)

	// compiledRegex, _ := regexp.Compile("^/bpm/.*")
	compiledRegex, _ := regexp.Compile(" ")
	uri := "/qtapi/base/bpm/process/instance/list"
	ok := compiledRegex.MatchString(uri)
	fmt.Println("ok:", ok)

	r, err := regexp.Compile("* * *") // error parsing regexp: missing argument to repetition operator: *
	r, err = regexp.Compile("*/.*")   // error parsing regexp: missing argument to repetition operator: *
	r, err = regexp.Compile("?/?")    // error parsing regexp: missing argument to repetition operator: *
	r, err = regexp.Compile("   ")    // error parsing regexp: missing argument to repetition operator: *
	r, err = regexp.Compile("无")

	r, err = regexp.Compile("?Add") // error parsing regexp: missing argument to repetition operator: `?`
	fmt.Println("r:", r)
	fmt.Println("err:", err)

}
