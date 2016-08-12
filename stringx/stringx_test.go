package stringx

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	s := ""
	empty := IsEmpty(s)

	t.Log(empty)
}
