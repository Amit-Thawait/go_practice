package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["Amit"] = 29
	ages["Manish"] = 23
	ages["John"] = 40
	fmt.Println("Map ages :", ages)

	delete(ages, "John")
	fmt.Println("Map ages after delete :", ages)

	months := map[string]int{
		"Jan": 1,
		"Feb": 2,
	}
	fmt.Println("Months :", months)

	_, ok := months["Abc"]
	fmt.Println("ok value :", ok)
}
