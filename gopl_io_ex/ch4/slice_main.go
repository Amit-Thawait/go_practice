package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)             // ["Apr" "May" "Jun"]
	fmt.Println(summer)         // ["Jun" "Jul" "Aug"]
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[Jun Jul Aug Sep Oct]"
}
