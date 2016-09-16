package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		s = s + sep + arg
	}
	fmt.Println(s)
}