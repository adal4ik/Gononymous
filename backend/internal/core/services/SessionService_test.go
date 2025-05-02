package services

import (
	"testing"
)

func TestPicker_Pick(t *testing.T) {
	p := NewPicker()
	hm := make(map[int]bool)

	for i := 1; i <= 827; i++ {
		id := p.Pick()
		if hm[id] {
			t.Error("Not unique id: ", id)
		}
		hm[id] = true
	}
}
