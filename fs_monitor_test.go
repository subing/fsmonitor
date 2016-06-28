package fsmonitor

import (
	"log"
	"testing"
	"time"
)

func echo() error {
	log.Print("echo")
	return nil
}

func maxPrint() error {
	log.Print("max.txt modifyed")
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
