package cmdrevive

import (
	"fmt"
	"github.com/go-fsnotify/fsnotify"
	"io"
	"log"
	"os/exec"
	"regexp"
	"sync"
)

func Start(targetDirs []string, pattern, cmdStr string, args []string) {

	targetFileNameRegex := regexp.MustCompile(pattern)

	var cmd *exec.Cmd = nil

	mutex := new(sync.Mutex)

	cmd = exec.Command(cmdStr, args...)
	go doEventTrigger(cmd)

	eventDriven(targetDirs, func(changedFileName string) {
		mutex.Lock()
		defer mutex.Unlock()
		if targetFileNameRegex.MatchString(changedFileName) {

			if cmd != nil {
				// Restart child process.
				if cmd.Process == nil {
					return
				}

				cmd.Process.Kill()
				cmd.Process.Wait()
			}

			cmd = exec.Command(cmdStr, args...)
			go doEventTrigger(cmd)
		}
	})
}

// eventDriven do something if changed file status.
func eventDriven(targetDirs []string, something func(changedFileNme string)) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				switch {
				case event.Op&fsnotify.Write == fsnotify.Write,
					event.Op&fsnotify.Create == fsnotify.Create,
					event.Op&fsnotify.Remove == fsnotify.Remove,
					event.Op&fsnotify.Rename == fsnotify.Rename:
					something(event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error: err := <-watcher.Errors:", err)
				done <- true
			}
		}
	}()

	for _, dir := range targetDirs {
		fmt.Println(dir)
		err = watcher.Add(dir)
		if err != nil {
			panic(err)
		}
	}

	<-done
}

func doEventTrigger(cmd *exec.Cmd) {

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("error: cmd.StdoutPipe(): ", err)
		return
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Println("error: cmd.StderrPipe(): ", err)
		return
	}

	defer stdoutPipe.Close()
	defer stderrPipe.Close()

	err = cmd.Start()
	if err != nil {
		log.Println("error: cmd.Start(): ", err)
		return
	}

	go printOutput(stdoutPipe)
	go printOutput(stderrPipe)

	cmd.Wait()
}

func printOutput(reader io.Reader) {
	var (
		err error
		n   int
	)
	buf := make([]byte, 1024)

	for {
		if n, err = reader.Read(buf); err != nil {
			break
		}

		fmt.Println(string(buf[0:n]))
	}

	if err != io.EOF {
		log.Println("error: err != io.EOF: " + err.Error())
	}
}
