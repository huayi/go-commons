package stringx

import (
	"strconv"
	"strings"
	"fmt"
)

func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func CharAt() {

}

func Matches() {

}

func SplitRmEmpty(s, sep string) []string {
	var arr []string
	for _, a := range strings.Split(s, sep) {
		if len(a) != 0 {
			arr = append(arr, a)
		}
	}
	return arr
}

// 截取字符串
// 建议使用go本身的语法 str[1:5]
func SubString(s string, start, end int) string {
	return s[start:end]
}

func ToCharArray() {

}

func String(i interface{}) string {

	switch s := i.(type) {
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case int:
		return strconv.FormatInt(int64(i.(int)), 10)
	case []byte:
		return string(s)
	case nil:
		return ""
	case fmt.Stringer:
		return s.String()
	case error:
		return s.Error()
	default:
		fmt.Errorf("Unable to Cast %#v to string", i)
		return ""
	}
}



