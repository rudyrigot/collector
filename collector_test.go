package collector

import "testing"

func TestC(t *testing.T) {
	assertEqual(t, len([]int{0, 1, 2}), len(C([]int{0, 1, 2})), "C function should return array of same length")
	assertIntArrayEqual(t, []int{0, 1, 2}, C([]int{0, 1, 2}), "C function should return same array")
}
