package main

import "fmt"

func main() {
	x := "HELLO"
	fmt.Printf("%v", 'a')
	fmt.Printf("%v", 'A')
	for _, x := range x {
		x = x + 32
		fmt.Printf("%c", x)
	}
	fmt.Println(x)
}
