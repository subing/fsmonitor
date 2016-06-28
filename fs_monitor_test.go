package fsmonitor

import (
	"fmt"
	"testing"
	"time"
)

func echo() error {
	fmt.Println("echo")
	return nil
}

func maxPrint() error {
	fmt.Println("max.txt modifyed")
	return nil
}
func TestWatch(t *testing.T) {
	AddWatch("./", "hello.txt", echo)
	AddWatch("./", "max.txt", maxPrint)
	Start()
	stop()
}

func stop() {
	time.Sleep(time.Second * 10)
	Stop()
}
