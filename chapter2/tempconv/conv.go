package tempconv

//CTof converts a Celsius temperature to Fahrenheit
func CTof(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k)
}

//KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 + 32)
}
