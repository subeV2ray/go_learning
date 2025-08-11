package array_test

import "testing"

func TestArrayTest(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3}
	arr2 := [...]int{1, 3, 4, 5}

	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3}
	for _, e := range arr {
		t.Log(e)
	}
	t.Log(arr[1:])
}
