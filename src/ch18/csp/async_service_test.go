package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("other task start")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("other task end")
}

//func TestService(t *testing.T) {
//	fmt.Println(service())
//	otherTask()
//}

func AsyncService() chan string {
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("async service start")
		retCh <- ret
		fmt.Println("async service end")
	}()

	return retCh
}

func TestAsyncService(t *testing.T) {
	a := AsyncService()
	otherTask()
	fmt.Println(<-a)
}
