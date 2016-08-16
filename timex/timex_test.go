package timex

import (
	"testing"
	"time"
	"fmt"
)

func TestTime(t *testing.T) {
	ti, err := time.Parse(time.RFC3339, "2015-12-12T12:12:11")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ti)
}