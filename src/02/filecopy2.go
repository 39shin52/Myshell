// 演習課題2

package main

import (
	"io"
	"os"
)

func main() {
	// open input file and make buffered reader
	fi, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// open output file and make buffered writer
	fo, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffered to read data
	buf := make([]byte, 1024)

	// copy whole content of
	// the input file to output file
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
}
