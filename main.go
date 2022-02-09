package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Log(x)
}

func main() {
	x := 2.0
	e := 1e-07
	fmt.Printf("backward: %v\n", backward(f, x, e))
	fmt.Printf("forward: %v\n", forward(f, x, e))
	fmt.Printf("central: %v\n", central(f, x, e))
	fmt.Printf("second: %v\n", second(f, x, e))

	ll := 1e-08 // lower limit
	ul := 1.0   // upper limit
	n := int64(100)
	fmt.Printf("sum: %v\n", integral(f, ll, ul, n))
}

func central(f func(x float64) float64, x float64, e float64) float64 {
	return (f(x+e) - f(x-e)) / (2 * e)
}

func backward(f func(x float64) float64, x float64, e float64) float64 {
	return (f(x) - f(x-e)) / e
}

func forward(f func(x float64) float64, x float64, e float64) float64 {
	return (f(x+e) - f(x)) / e
}

func second(f func(x float64) float64, x float64, e float64) float64 {
	return (f(x+e) - 2*f(x) + f(x-e)) / (e * e)
}

// Simpson's 1/3 Rule
func integral(f func(x float64) float64, lower float64, upper float64, steps int64) float64 {
	step := (upper - lower) / float64(steps)
	sum := f(lower) + f(upper)
	for i := int64(1); i <= steps-1; i++ {
		k := lower + float64(i)*step
		if i%2 == 0 {
			sum = sum + 2*(f(k))
		} else {
			sum = sum + 4*(f(k))
		}
	}
	sum = sum * step / 3
	return sum
}
