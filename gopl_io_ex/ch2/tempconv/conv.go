package tempconv

func CToF(c Celsius) Farenhiet {
	return Farenhiet((c*9)/5 + 32)
}

func FToC(f Farenhiet) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
