package main

import "fmt"

func f(x float64) float64 {
	return x * x
}

func main() {
	x := 0.0    // точка, в которой вычисляем производную
	e := 0.0001 // шаг, с которым вычисляем производную

	// приближенно вычисляем первую производную различными способами
	fl := (f(x) - f(x-e)) / e         // левая
	fr := (f(x+e) - f(x)) / e         // правая
	fc := (f(x+e) - f(x-e)) / (2 * e) // центральная

	// приближенно вычисляем вторую производную
	f2 := (f(x+e) - 2*f(x) + f(x-e)) / (e * e)
	fmt.Println("fl:", fl)
	fmt.Println("fr:", fr)
	fmt.Println("fc:", fc)
	fmt.Println("f2:", f2)
}
