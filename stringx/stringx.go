package stringx

import (
	"strconv"
	"strings"
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

func SubString() {

}

func ToCharArray() {

}

func ValueOfInt(i int) string {
	return strconv.Itoa(i)
}


func ValueOfFloat32(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

func ValueOfFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}



