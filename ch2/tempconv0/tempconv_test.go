package tempconv

import "fmt"

func example1() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch
}

func example2() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)  // "true"
	fmt.Println(f >= 0)  // "true"
	fmt.Println(c == f)  // compile error: type mismatch
	fmt.Println(c == Celsius(f))  // "true"!
}

func example3() {
	c := FToC(212.0)
	fmt.Println(c.String())  // "100°C"
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
