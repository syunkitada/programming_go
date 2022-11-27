package api

import "testing"

func TestGetResult(t *testing.T) {
	result := GetResult(1, 2)
	if result != 3 {
		t.Errorf("actual %d\nwant %d", result, 3)
	}
}
