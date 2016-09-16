package tempconv

import "fmt"

type Celsius float64
type Farenhiet float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g deg C", c)
}

func (f Farenhiet) String() string {
	return fmt.Sprintf("%g def F", f)
}
