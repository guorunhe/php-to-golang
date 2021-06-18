package string

/**
说明
将字符串转换为数组

参数
string
输入字符串。

split_length
每一段的长度。

返回值
如果指定了可选的 split_length 参数，返回数组中的每个元素均为一个长度为 split_length 的字符块，否则每个字符块为单个字符。

如果 split_length 小于 1，返回 false。如果 split_length 参数超过了 string 超过了字符串 string 的长度，整个字符串将作为数组仅有的一个元素返回。

*/
func StrSplit(s string, split_length int) []string {
	if split_length < 1 || split_length > len(s) {
		return []string{s}
	}
	res := make([]string, 0)

	i := 0
	for ; i+split_length < len(s); i += split_length {
		res = append(res, s[i:i+split_length])
	}

	if i != len(s) {
		res = append(res, s[i:len(s)])
	}
	return res
}
