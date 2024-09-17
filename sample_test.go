package main

import "testing"

func TestMult(t *testing.T) {
	var a, b, res = 2, 3, 6
	realResult := mult(a, b)
	if realResult != res {
		t.Errorf("%d != %d", realResult, res)
	}
}
