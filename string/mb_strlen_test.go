package string

import (
        "testing";
        "fmt"
)

func TestStrlen(t *testing.T) {
        fmt.Println("-------------------strlen test start--------------")
        s := "123,233,2332"
        strLen := MbStrlen(s)
        fmt.Printf("数字数组: %s 的长度是: %d", s, strLen)
        fmt.Println();

        s = "你好，世界哈哈"
        strLen = MbStrlen(s)
        fmt.Printf("汉字数组: %s 的长度是: %d", s, strLen)
        fmt.Println();
        t.Log(strLen)
        fmt.Println("-------------------strlen test end--------------")
}
