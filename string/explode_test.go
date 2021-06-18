package string

import (
	"fmt"
	"testing"
)

func TestExplode(t *testing.T) {
	fmt.Println("-------------------explode test start---------------")
	strArray := Explode(',', "123,233,2332", -1)
	fmt.Println("limit 是负数")
	for i, v := range strArray {
		fmt.Println(i, v)
	}
	fmt.Println("limit正数")
	strArray = Explode(',', "123,233,2332", 2)
	for _, v := range strArray {
		fmt.Println(v)
	}
	fmt.Println("limit大于个数")
	strArray = Explode(' ', "guorunhe study golang hahahahaha", 100)
	for _, v := range strArray {
		fmt.Println(v)
	}
	t.Log(strArray)
	fmt.Println("-------------------explode test end-----------------")
}
