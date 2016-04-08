package collector

import "testing"

func TestC(t *testing.T) {
	assertEqualM(t, len([]int{0, 1, 2}), len(C([]int{0, 1, 2})), "C function should return array of same length")
}
