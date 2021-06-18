package string

import (
	"fmt"
	"testing"
)

func TestStrSplit(t *testing.T) {
	fmt.Println("-------------------str_split test start--------------")
	strArray := StrSplit("123,233,2332", 0)
	fmt.Println("limit 是0")
	for i, v := range strArray {
		fmt.Println(i, v)
	}
	fmt.Println("limit 小于字符串长度的正数")
	strArray = StrSplit("123,233,2332", 5)
	for _, v := range strArray {
		fmt.Println(v)
	}
	fmt.Println("limit大于字符串长度的正数")
	strArray = StrSplit("guorunhe study golang hahahahaha", 100)
	for _, v := range strArray {
		fmt.Println(v)
	}
	t.Log(strArray)
	fmt.Println("-------------------str_split test end--------------")
}
