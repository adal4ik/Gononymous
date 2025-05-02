package test

import (
	"Gononymous/internal/core/services"
	"testing"
)

func TestPicker_Pick(t *testing.T) {
	p := services.NewPicker()
	hm := make(map[int]bool)

	for i := 1; i <= 827; i++ {
		id := p.Pick()
		if hm[id] {
			t.Error("Not unique id: ", id)
		}
		hm[id] = true
	}
}
