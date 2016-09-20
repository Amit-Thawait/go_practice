package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d	cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	a := []int{1}
	a = appendSlice(a, 2, 3)
	fmt.Println("a after appending 2, 3 :", a)
	a = appendSlice(a, 4, 5)
	fmt.Println("a after appending 4, 5 :", a)

	// In-built append function
	fmt.Println("a after appending 6, 7, 8 using append :", append(a, 6, 7, 8))
}

func appendSlice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		fmt.Println("=====zcap=%d====zlen=%d", zcap, zlen)
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
