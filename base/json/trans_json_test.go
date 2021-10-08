package main

import (
	"encoding/json"
	"fmt"
	"gitlab.pri.ibanyu.com/quality/dry.git/errors"
	"strconv"
	"testing"
)

//测试空字符、nil Json转码
func TestTransJson(t *testing.T) {

	firstOrderId, err := strconv.ParseInt("", 10, 64)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(firstOrderId)

	TransInt64ArrToJson(nil)
	TransInt64ArrToJson([]int64{})

	TransJsonToInt64Arr("")
	TransJsonToStringArr("")
	TransJsonToStringArr("[]")
}

func TransInt64ArrToJson(data []int64) string {
	op := errors.Op("TransInt64ArrToJson")
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("%s err %v \n", op, err)
		return ""
	}
	fmt.Printf("%s res %s \n", op, string(res))
	return string(res)
}

func TransJsonToInt64Arr(jsonStr string) (result []int64) {
	op := errors.Op("TransJsonToInt64Arr")
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Printf("%s err %v \n", op, err)
		return nil
	}
	fmt.Printf("%s res %v \n", op, result)
	return result
}

func TransJsonToStringArr(jsonStr string) (result []string) {
	op := errors.Op("TransJsonToStringArr")
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Printf("%s err %v \n", op, err)
		return nil
	}
	fmt.Printf("%s res %v \n", op, result)
	return result
}
