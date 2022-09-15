package main

import "fmt"

func check(s int) int {
	count := 0
	for i := 1; i < s; i++ {
		if s%i == 0 {
			count++
		}
	}
	if count == 2 {
		return 1
	} else {
		return 0
	}
}

func main() {
	n := 859433
	count := 0

	for i := 1; i < n; i++ {
		if check(i) == 1 {
			count++
		}
	}
	if count > 0 {
		fmt.Printf("%d は素数です。%d 番目です。\n", n, count)
	} else {
		fmt.Printf("%d は素数ではありません。\n", n)
	}
}
