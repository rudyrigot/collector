package collector

import "testing"

func TestC(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{len([]int{0, 1, 2}), len(C([]int{0, 1, 2}))},
	}
	for _, c := range cases {
		got := c.in
		if got != c.want {
			t.Errorf("C: %q != %q", c.in, c.want)
		}
	}
}
