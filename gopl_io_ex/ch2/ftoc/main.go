package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("Freezing Point = %gF or %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("Boiling Point = %gF or %gC\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
