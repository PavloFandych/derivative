package main

import (
	"fmt"
	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/integrate/quad"
	"math"
)

func f(x float64) float64 {
	return math.Log(x)
}

func main() {
	//derivative
	x := 2.0
	e := 1e-07
	fmt.Printf("backward: %v\n", backward(f, x, e))
	fmt.Printf("forward: %v\n", forward(f, x, e))
	fmt.Printf("central: %v\n", central(f, x, e))
	fmt.Printf("second: %v\n", second(f, x, e))
	//https://pkg.go.dev/gonum.org/v1/gonum
	fmt.Println("f'(0) ≈", fd.Derivative(f, x, nil))
	df := fd.Derivative(f, x, &fd.Settings{
		Formula: fd.Forward,
		Step:    e,
	})
	fmt.Println("f'(0) ≈", df)

	//integral
	ll := 1e-08 // lower limit
	ul := 1.0   // upper limit
	n := int64(100)
	fmt.Printf("sum: %v\n", integral(f, ll, ul, n))
	//https://pkg.go.dev/gonum.org/v1/gonum
	fmt.Printf("sum gonum = %v\n", quad.Fixed(f, ll, ul, 100, nil, 0))

	//newton
	fmt.Printf("newton: %v\n", newtonRaphson(f, 0.5))
}

func central(f func(float64) float64, x float64, e float64) float64 {
	return (f(x+e) - f(x-e)) / (2 * e)
}

func backward(f func(float64) float64, x float64, e float64) float64 {
	return (f(x) - f(x-e)) / e
}

func forward(f func(float64) float64, x float64, e float64) float64 {
	return (f(x+e) - f(x)) / e
}

func second(f func(float64) float64, x float64, e float64) float64 {
	return (f(x+e) - 2*f(x) + f(x-e)) / (e * e)
}

// Simpson's 1/3 Rule
func integral(f func(float64) float64, lower float64, upper float64, steps int64) float64 {
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

// x - initial guess
func newtonRaphson(f func(float64) float64, x float64) float64 {
	h := f(x) / central(f, x, 1e-07)
	for math.Abs(h) >= 0.0001 {
		h = f(x) / central(f, x, 1e-07)
		x = x - h
	}
	return x
}
