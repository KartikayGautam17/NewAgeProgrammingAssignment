package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")

	a := 15
	b := 4
	x := 7.5
	y := 2.5

	fmt.Println("Integers:", a, "and", b)
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)

	fmt.Println("Floats:", x, "and", y)
	fmt.Println("x + y =", x+y)
	fmt.Println("x - y =", x-y)
	fmt.Println("x * y =", x*y)
}
