// 課題1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
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

			command := strings.Split(strings.TrimSpace(scanner.Text()), " ")

			cpath, err := exec.LookPath(command[0])
			if err != nil {
				log.Fatalf("%s not found in $PATH.", scanner.Text())
			}

			args := command[0:]
			attr := syscall.ProcAttr{
				Files: []uintptr{0, 1, 2},
			}

			pid, err := syscall.ForkExec(cpath, args, &attr)
			if err != nil {
				panic(err)
			}

			proc, err := os.FindProcess(pid)
			if err != nil {
				panic(err)
			}

			status, err := proc.Wait()
			if err != nil {
				panic(err)
			}

			if !status.Success() {
				fmt.Println(status.String())
			}
		}
		fmt.Printf("./Myshell[%02d]> ", count)
	}
}