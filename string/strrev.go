package string

// import (
//         "reflect";
//         "fmt"
// )
/**
反转字符串

说明 ¶
strrev(string $string): string
返回 string 反转后的字符串。

参数 ¶
string
待反转的原始字符串。

返回值 ¶
返回反转后的字符串。
*/
/**
遇到的问题：
1.数组的值变更操作（不允许）
        采用slice来操作.
*/

func Strrev(s string) string {
	s1 := []rune(s)
	i, j := 0, len(s1)-1
	// sl := make([]rune, len(s))
	// fmt.Print(reflect.TypeOf(sl))
	for i < j {
		s1[i], s1[j] = s1[j], s1[i]
		// tmp := s1[j]
		// sl[j] = s1[i]
		// sl[i] = tmp
		i++
		j--
	}
	// fmt.Print(sl)

	return string(s1)
}
