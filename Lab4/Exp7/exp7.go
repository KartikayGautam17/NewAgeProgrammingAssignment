package main

import (
	"fmt"
	"math"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %g by zero", a)
	}
	return a / b, nil
}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot take square root of negative number %g", x)
	}
	return math.Sqrt(x), nil
}

func Average(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("cannot find average of an empty list")
	}
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums)), nil
}

func main() {
	if result, err := Divide(10, 4); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 4 =", result)
	}

	if _, err := Divide(10, 0); err != nil {
		fmt.Println("Error:", err)
	}

	if result, err := SquareRoot(49); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("sqrt(49) =", result)
	}

	if _, err := SquareRoot(-9); err != nil {
		fmt.Println("Error:", err)
	}

	if result, err := Average([]float64{4, 8, 15, 16}); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average =", result)
	}

	if _, err := Average([]float64{}); err != nil {
		fmt.Println("Error:", err)
	}
}
