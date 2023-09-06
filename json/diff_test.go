package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bytedance/sonic/encoder"
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

// 标准库中默认会开启html Escape，而Sonic 出于性能损耗默认不开启
func TestEncode(t *testing.T) {
	data := map[string]string{"&&": "<>"}

	var w1 = bytes.NewBuffer(nil)
	enc1 := json.NewEncoder(w1)
	//enc1.SetEscapeHTML(true)
	err := enc1.Encode(data)
	assert.NoError(t, err)

	var w2 = bytes.NewBuffer(nil)
	enc2 := encoder.NewStreamEncoder(w2)
	//enc2.SetEscapeHTML(true)
	err = enc2.Encode(data)
	assert.NoError(t, err)
	fmt.Printf("%v%v", w1.String(), w2.String())
}

type User struct {
	ID   int64  `json:"id`
	Name string `json:"name`
	Sex  string `json:"sex`
}

/*
JSON Umarshal 针对同一个对象是增量更新的，而是全量更新
*/
func TestStdJsonUnmarshal(t *testing.T) {
	user1 := User{}
	var oldData = `{"ID":1,"Name":"A","Sex":""}`
	err := json.Unmarshal([]byte(oldData), &user1)
	if err != nil {
		return
	}
	var newData = `{"Name":"B","Sex":"man"}`
	err = json.Unmarshal([]byte(newData), &user1)
	if err != nil {
		return
	}
	data, err := json.Marshal(user1)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	// 结果：{"ID":1,"Name":"B","Sex":"man"}
}
