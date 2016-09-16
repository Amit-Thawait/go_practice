package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	fmt.Printf("Type of months is : %T\n", months)
	fmt.Println(months[0])
	fmt.Println(months)
	fmt.Println(months[3])

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Printf("Type of days is : %T\n", days)
	fmt.Println(days)
	fmt.Println(days[5])

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "E", GBP: " ! ", RMB: "R"}
	fmt.Println(RMB, symbol[RMB])
}
