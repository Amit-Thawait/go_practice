package tempconv

type Celsius float64
type Farenhiet float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func FToC(f Farenhiet) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Farenhiet {
	return Farenhiet((c*9)/5 + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g deg C", c)
}
