package main

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = iota << 1
	Writable
	Executable
)

func TestConstTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry(t *testing.T) {
	a := 6 // 0110

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
