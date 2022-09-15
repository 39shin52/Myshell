// 課題4(任意)

package main

import (
	"fmt"
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

	size, err := os.Stat(os.Args[1])
	if err != nil {
		panic(err)
	}

	b, err := syscall.Mmap(fi, 0, int(size.Size()), syscall.PROT_READ, syscall.MAP_SHARED)

	for i, v := range b {
		if (v >= 48 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122) || v == 32 || v == 10 {
			if _, err := syscall.Write(fo, b[i:i+1]); err != nil {
				panic(err)
			}
		}
	}

	stat, err := os.Stat(os.Args[2])
	if err != nil {
		panic(err)
	}

	fmt.Println(stat.Size())

}
