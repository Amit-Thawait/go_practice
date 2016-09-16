package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[6])
	fmt.Printf("%c %c\n", s[0], s[6])
}
