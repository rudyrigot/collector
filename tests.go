package collector

import "testing"

func assertEqual(t *testing.T, in interface{}, want interface{}, message string) {
	if in != want {
		t.Errorf("[%#v != %#v] %s", in, want, message)
	}
}

func assertIntArrayEqual(t *testing.T, in []int, want []int, message string) {
	if len(in) != len(want) {
		t.Errorf("[%#v != %#v] %s", in, want, message)
	} else {
		for i, elem := range in {
			if want[i] != elem {
				t.Errorf("[%#v != %#v] %s", in, want, message)
				return
			}
		}
	}
}
