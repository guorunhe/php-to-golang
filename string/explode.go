package string

/**
说明
使用一个字符串分割另一个字符串

参数
separator
边界上的分隔字符。

string
输入的字符串。

limit
如果设置了 limit 参数并且是正数，则返回的数组包含最多 limit 个元素，而最后那个元素将包含 string 的剩余部分。

如果 limit 参数是负数，则返回除了最后的 -limit 个元素外的所有元素。

如果 limit 是 0，则会被当做 1。

返回值
此函数返回由字符串组成的 array，每个元素都是 string 的一个子串，它们被字符串 separator 作为边界点分割出来。
*/
func Explode(separator rune, s string, limit int) []string {
	if limit == 0 {
		limit = 1
	}
	tag := false
	if limit < 0 {
		// limit小于零
		limit = -limit
		tag = true
	}
	res := make([]string, 0)
	j, count := 0, 0
	for i, c := range s {
		if !tag && count < limit-1 && separator == c {
			res = append(res, s[j:i])
			j = i + 1
			count++
		} else if tag && separator == c {
			res = append(res, s[j:i])
			j = i + 1
			count++
		}
	}
	res = append(res, s[j:len(s)])

	if tag {
		res = res[:len(res)-limit]
	}
	return res
}
