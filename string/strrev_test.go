package string

import (
	"fmt"
	"testing"
)

func TestStrrev(t *testing.T) {
	fmt.Println("-------------------strrev test start--------------")
	str := Strrev("guorunhe")
	fmt.Println(str)
	t.Log(str)
	fmt.Println("-------------------strrev test end--------------")
}
