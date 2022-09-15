// 課題2

package main

import (
	"io"
	"os"
	"syscall"
)

func main() {
	// open input file and make a buffered reader
	fi, err := syscall.Open(os.Args[1], syscall.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := syscall.Close(fi); err != nil {
			panic(err)
		}
	}()

	// open output file and make a buffered writer
	fo, err := syscall.Open(os.Args[2], syscall.O_WRONLY|syscall.O_CREAT|syscall.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := syscall.Close(fo); err != nil {
			panic(err)
		}
	}()

	// make buffer to read data
	buf := make([]byte, 1024)

	// copy the whole data of
	// the input file to output file
	for {
		n, err := syscall.Read(fi, buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			if (buf[i] >= 48 && buf[i] <= 57) || (buf[i] >= 65 && buf[i] <= 90) || (buf[i] >= 97 && buf[i] <= 122) || buf[i] == 32 || buf[i] == 10 {
				if _, err := syscall.Write(fo, buf[i:i+1]); err != nil {
					panic(err)
				}
			}
		}
	}
}
