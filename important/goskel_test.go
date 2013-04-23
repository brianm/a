package important

import (
	"testing"
)

func Test_Math(t *testing.T) {
	r := 2 + 2
	if r != 4 {
		t.Errorf("int math is broken, 2 + 2 = '%d' !?", r)
	}
}
