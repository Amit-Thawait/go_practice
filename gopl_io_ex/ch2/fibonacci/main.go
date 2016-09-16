package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(fib(num))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 1; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
