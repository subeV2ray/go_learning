package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 8 // 111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
