// 課題2 rawsyscall版

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func splitText(text string) {
	if strings.Contains(text, "?") && strings.Contains(text, ":") {
		beforeCol := strings.Split(text, ":")[0]

		firstCom := strings.Split(beforeCol, "?")[0]
		if doCommand(firstCom) == 0 {
			doCommand(strings.Split(beforeCol, "?")[1])
		} else {
			doCommand(strings.Split(text, ":")[1])
		}
	} else {
		doCommand(text)
	}
}

func doCommand(text string) int {
	command := strings.Split(strings.TrimSpace(text), " ")
	cpath, err := exec.LookPath(command[0])
	if err != nil {
		log.Printf("%s not found in $PATH.", text)
		return 1 // ここでreturnすることでforcexecにcpathがいかないのでpanicまで到達しない
	}

	args := command[0:]
	env := os.Environ()

	if err := syscall.Exec(cpath, args, env); err != nil {
		panic(err)
	}

	return 0
}

func main() {
	pid, r2, _ := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)

	// fixup fot macOS
	if runtime.GOOS == "darwin" && r2 == 1 {
		pid = 0
	}

	if pid > 0 {
		// parent process
		var rusage syscall.Rusage
		status := syscall.WaitStatus(0)
		_, err := syscall.Wait4(int(pid), &status, syscall.WSTOPPED, &rusage)
		if err != nil {
			panic(err)
		}

		if status != 0 {
			fmt.Println("exit status", status)
		}

		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	fmt.Print("./Myshell[00]> ")

	for scanner.Scan() {
		if scanner.Text() == "bye" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		if len(scanner.Text()) > 0 {
			count++
			// doCommand(scanner.Text())
			splitText(scanner.Text())
		}

		fmt.Printf("./Myshell[%02d]> ", count)
	}
}
