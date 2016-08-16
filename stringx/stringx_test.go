package stringx

import (
	"testing"
	"strconv"
	"strings"
)

func TestIsEmpty(t *testing.T) {
	s := ""
	empty := IsEmpty(s)

	t.Log(empty)
}

func TestValueOfFloat(t *testing.T) {
	var a float32
	a = 1.23
	f1 := ValueOfFloat32(a)
	f2 := ValueOfFloat64(1.2)
	t.Log(f1)
	t.Log(f2)
	strconv.FormatFloat(float64(a), 'f', -1, 64)

	s := "1,,2,3,"
	r := strings.SplitAfter(s, ",")
	t.Log(r)

}