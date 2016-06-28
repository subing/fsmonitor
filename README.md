# fsmonitor

detect file's modification and execute callback function when file modified

**useage:**

	go get github.com:subing/fsmonitor.git


```go

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
```
```
