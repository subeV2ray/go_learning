package loop_test_go

import (
	"testing"
)

// loop
func TestWhileLoop(t *testing.T) {
	n := 0

	for n < 5 {
		t.Log(n)
		n++
	}
}

// condition
func TestCondition(t *testing.T) {

	if v, err := someFun(); err == nil {
		t.Log(v)
	} else {
		t.Log("1 == 1")
	}
}

func someFun() (interface{}, interface{}) {
	return "1", nil
}
