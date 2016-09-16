package main

import (
	"fmt"
	"github.com/amit-thawait/gopl_io_ex/ch2/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.FreezingC)
	fmt.Printf("In formatted %%g output : %g\n", tempconv.BoilingC)
	fmt.Printf("In formatted %%v output : %v\n", tempconv.BoilingC)
}
