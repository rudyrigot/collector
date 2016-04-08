package collector

import "testing"

func TestC(t *testing.T) {
	init := []string{"foo", "bar", "foo"}
	assertEqual(t, len(init), len(CString(init)), "C function should return array of same length")
	assertStringArrayEqual(t, init, CString(init), "C function should return same array")
}
