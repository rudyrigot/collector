package collector

import "testing"

func assertEqual(t *testing.T, in interface{}, want interface{}) {
	if in != want {
		t.Errorf("[%#v != %#v]", in, want)
	}
}

func assertEqualM(t *testing.T, in interface{}, want interface{}, message string) {
	if in != want {
		t.Errorf("[%#v != %#v] %s", in, want, message)
	}
}
