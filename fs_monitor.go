package fsmonitor

import (
	"github.com/howeyc/fsnotify"
	"log"
	"path/filepath"
)

type Func func() error

type FsMonitor struct {
	path    []string
	fileOps map[string]Func
}

var watcher *fsnotify.Watcher

var fsMonitor FsMonitor

func init() {
	fsMonitor.fileOps = make(map[string]Func)
}
func AddWatch(path, file string, ops Func) {
	fsMonitor.path = append(fsMonitor.path, path)
	fsMonitor.fileOps[file] = ops
}

func DeleteWatch(path, file string) {
	watcher.RemoveWatch(path)
}

var done = make(chan bool)

func Start() error {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Print("event:", ev)
				if ev.IsModify() && ev.IsAttrib() {
					ops := fsMonitor.fileOps[filepath.Base(ev.Name)]
					if ops != nil {
						ops()
					}
				}
			case err := <-watcher.Error:
				log.Print("error:", err)
			case <-done:
				watcher.Close()
				return
			}
		}
	}()
	for _, v := range fsMonitor.path {
		log.Print("path: ", v)
		err = watcher.Watch(v)
		if err != nil {
			log.Print(err.Error())
			return err
		}
	}
	return nil
}

func Stop() {
	done <- true
}
