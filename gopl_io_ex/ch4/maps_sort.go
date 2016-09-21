package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"Amit":   29,
		"Manish": 23,
		"John":   40,
	}

	fmt.Println("Input Map :", ages)

	// var names []string // Not efficient
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
