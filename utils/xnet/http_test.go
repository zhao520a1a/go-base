package xnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type AddCacheData struct {
	Version       int             `json:"version"`
	Encoding      string          `json:"encoding"`
	Command       string          `json:"command"`
	Namespace     string          `json:"namespace"`
	Prefix        string          `json:"prefix"`
	PayloadFields []string        `json:"payload_fields"`
	PayloadData   [][]interface{} `json:"payload_data"`
}

func TestHttp(t *testing.T) {
	paramMap := make(map[string]interface{})

	url := "fddf"
	paramMap["version"] = 1
	paramMap["task_id"] = 1
	paramMap["encoding"] = "command"
	paramMap["command"] = "SET"
	paramMap["namespace"] = "test/test"
	paramMap["prefix"] = "indexer"

	var payloadData [][]interface{}

	for i := 0; i < 100; i++ {
		item := []interface{}{
			"k" + strconv.Itoa(i),
			"v" + strconv.Itoa(i),
			30,
		}
		payloadData = append(payloadData, item)
	}
	paramMap["payload_data"] = payloadData

	SendRequest(url, paramMap, "")
	//Post(url,paramMap,"application/json")
}

func SendRequest(url string, paramMap map[string]interface{}, routerGroup string) {
	jsonStr, _ := json.Marshal(paramMap)
	payload := bytes.NewBuffer(jsonStr)
	req, _ := http.NewRequest("POST", url, payload)
	if len(routerGroup) > 0 {
		req.Header.Add("ipalfish-group", routerGroup)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "")
	req.Header.Add("Cookie", "")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
