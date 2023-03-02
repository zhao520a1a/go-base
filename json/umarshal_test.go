package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	/*
		JSON 解析不仅支持是对象、数组，同时也可以是字符串、数值、布尔值以及空值
		注意:字符串中的双引号不能缺，否则将不是一个合法的 JSON 序列，会返回错误。
	*/
	var s string
	err := json.Unmarshal([]byte(`"Hello, world!"`), &s)
	assert.NoError(t, err)
	fmt.Println(s)

	/*
		如果遇到大小写问题，会尽可能地进行大小写转换。
		一个 key 与结构体中的定义不同，但忽略大小写后是相同的，那么依然能够为字段赋值
	*/
	cert := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err = json.Unmarshal([]byte(`{"UserName":"root","passWord":"123456"}`), &cert)
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println("username =", cert.Username)
		fmt.Println("password =", cert.Password)
	}
}
