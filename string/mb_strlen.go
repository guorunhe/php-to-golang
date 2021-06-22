package string

/**
获取字符串的长度

说明 ¶
mb_strlen(string $str, string $encoding = mb_internal_encoding()): mixed
获取一个 string 的长度。

参数 ¶
str
要检查长度的字符串。

encoding
encoding 参数为字符编码。如果省略或是 null，则使用内部字符编码。

返回值 ¶
返回具有 encoding 编码的字符串 str 包含的字符数。 多字节的字符被计为 1。

如果给定的 encoding 无效则返回 false。
*/
func MbStrlen(s string) int {
        // []rune 和 []byte 区别
        // https://learnku.com/articles/23411/the-difference-between-rune-and-byte-of-go
        return len([]rune(s))
}