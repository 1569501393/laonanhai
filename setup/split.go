package setup

import (
	"strings"
)

// Split 定义一个分割字符串的函数
// 输入 a:b:c 和 : 输出 [a,b,c]
func Split(str, sep string) (result []string) {
	idx := strings.Index(str, sep)

	for idx >= 0 {
		result = append(result, str[:idx])
		str = str[idx+len(sep):]
		idx = strings.Index(str, sep)
	}
	result = append(result, str)
	return
}
