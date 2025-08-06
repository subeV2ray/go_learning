package test

import (
	"testing"
)

func TestFirstTry(t *testing.T) {
	//var (
	//	a = 1
	//	b = 2
	//)

	a := 1
	b := 2

	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log()
}
