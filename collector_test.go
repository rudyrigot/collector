package collector

import (
	"fmt"
	"testing"
)

func TestC(t *testing.T) {
	init := []string{"foo", "bar", "foo"}
	want := operationSet{initialArray: []string{"foo", "bar", "foo"}}
	assertStringArrayEqual(t, CString(init).initialArray, want.initialArray, "CString result should have proper initialArray")
	assertEqual(t, CString(init).alreadyPerformed, false, "CString result should have alreadyPerformed at false")
	assertEqual(t, len(CString(init).ops), 0, "CString result should have empty ops")
}

func TestTransform(t *testing.T) {
	f := func(...string) string { return "hello" }
	opSet := operationSet{ops: []*operation{&operation{}}}
	in := opSet.Transform(f)
	assertEqual(t, len(in.ops), 2, "Transform should append a new operation")
	assertEqual(t, in.ops[1].op, transformOp, "Transform should append a new operation of the transformOp type")
}

func TestTransformInto(t *testing.T) {
	init := []string{"foo", "bar", "foo"}
	want := []string{"foofoo", "barbar", "foofoo"}
	f := func(s ...string) string { return fmt.Sprintf("%s%s", s, s) }
	var in []string
	err := CString(init).Transform(f).Into(&in)
	if err != nil {
		t.Errorf("Error happened while running Into: %s", err)
		return
	}
	assertStringArrayEqual(t, in, want, "Into should execute the Tranform operation")
}
