package stringx

import (
	"strconv"
	"testing"
	//"strings"
)

func TestIsEmpty(t *testing.T) {
	s := ""
	empty := IsEmpty(s)

	t.Log(empty)
}

func TestSplitRmEmpty(t *testing.T) {
	s := "1,2,3,"

	//r := strings.SplitAfter(s, ",")
	//t.Log(r)

	//r := strings.Split(s, ",")
	//t.Log(r)

	//s = "1,2,3,"
	r := SplitRmEmpty(s, ",")
	t.Log(r)
}

func TestSubString(t *testing.T) {
	s := "1234567"
	s1 := s[:3]
	t.Log(s1)

	s2 := SubString(s, 0, 3)
	t.Log(s2)

	//var s3 string
	t.Log(s2[0])
	t.Log(s2[1])
}

func TestValueOfFloat(t *testing.T) {
	var a float32
	a = 1.23
	f1 := ValueOfFloat32(a)
	f2 := ValueOfFloat64(1.2)
	t.Log(f1)
	t.Log(f2)
	strconv.FormatFloat(float64(a), 'f', -1, 64)

}

func TestString(t *testing.T) {
	type I int
	var i I
	i = 1
	//i := "a"
	//i := 1.25
	//i := "a"
	//i := "a"
	//var i interface{}
	s := String(i)
	t.Log(s)
}
