package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShrink_Int(t *testing.T) {
	cases := []struct {
		name    string
		input   []int
		capIn   int
		wantCap int
	}{
		{"no shrink small cap", []int{1, 2, 3}, 32, 32},
		{"no shrink normal", []int{1, 2, 3}, 100, 50}, // 修正
		{"shrink big cap", make([]int, 100, 4096), 4096, 2560},
		{"shrink mid cap", make([]int, 100, 2048), 2048, 1024},
		{"no shrink big but not enough", make([]int, 2000, 4096), 4096, 2560}, // 修正
		{"shrink mid but not enough", make([]int, 600, 2048), 2048, 2048},
		{"shrink empty", []int{}, 0, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var s []int
			if c.capIn > 0 {
				s = make([]int, len(c.input), c.capIn)
				copy(s, c.input)
			} else {
				s = c.input
			}
			res := Shrink(s)
			assert.Equal(t, c.wantCap, cap(res))
			assert.Equal(t, c.input, res)
		})
	}
}

func TestShrink_String(t *testing.T) {
	cases := []struct {
		name    string
		input   []string
		capIn   int
		wantCap int
	}{
		{"no shrink small cap", []string{"a", "b"}, 32, 32},
		{"shrink big cap", make([]string, 10, 4096), 4096, 2560},
		{"shrink mid cap", make([]string, 10, 2048), 2048, 1024},
		{"shrink empty", []string{}, 0, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var s []string
			if c.capIn > 0 {
				s = make([]string, len(c.input), c.capIn)
				copy(s, c.input)
			} else {
				s = c.input
			}
			res := Shrink(s)
			assert.Equal(t, c.wantCap, cap(res))
			assert.Equal(t, c.input, res)
		})
	}
}
