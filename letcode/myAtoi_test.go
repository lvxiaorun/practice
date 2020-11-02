package letcode

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	args := []string{"9223372036854775808"}
	for _, item := range args {
		t.Log(MyAtoi(item))
	}
}
