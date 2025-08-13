package ch8

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 无参，返回多个值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数作为参数，返回值为一个int值的函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start))
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFun(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
