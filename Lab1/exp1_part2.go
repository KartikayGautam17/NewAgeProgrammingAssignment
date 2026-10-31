package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const limit = 1000000

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) (string, bool) {
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil && text == "" {
		return "", false
	}
	return strings.TrimSpace(text), true
}

func readInt(prompt string) (int, bool) {
	for {
		text, ok := readLine(prompt)
		if !ok {
			return 0, false
		}
		n, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("That is not a whole number, try again.")
			continue
		}
		if n < -limit || n > limit {
			fmt.Printf("Number must be between %d and %d.\n", -limit, limit)
			continue
		}
		return n, true
	}
}

func readFloat(prompt string) (float64, bool) {
	for {
		text, ok := readLine(prompt)
		if !ok {
			return 0, false
		}
		f, err := strconv.ParseFloat(text, 64)
		if err != nil || math.IsNaN(f) || math.IsInf(f, 0) {
			fmt.Println("That is not a valid number, try again.")
			continue
		}
		if f < -limit || f > limit {
			fmt.Printf("Number must be between %d and %d.\n", -limit, limit)
			continue
		}
		return f, true
	}
}

func intCalc() bool {
	a, ok := readInt("Enter first integer: ")
	if !ok {
		return false
	}
	b, ok := readInt("Enter second integer: ")
	if !ok {
		return false
	}
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	return true
}

func floatCalc() bool {
	a, ok := readFloat("Enter first number: ")
	if !ok {
		return false
	}
	b, ok := readFloat("Enter second number: ")
	if !ok {
		return false
	}
	fmt.Printf("%g + %g = %g\n", a, b, a+b)
	fmt.Printf("%g - %g = %g\n", a, b, a-b)
	fmt.Printf("%g * %g = %g\n", a, b, a*b)
	return true
}

func main() {
	fmt.Println("Simple Calculator")

	for {
		fmt.Println()
		fmt.Println("1. Integer calculation")
		fmt.Println("2. Floating-point calculation")
		fmt.Println("3. Quit")

		choice, ok := readLine("Choose an option: ")
		if !ok {
			break
		}

		if choice == "1" {
			if !intCalc() {
				break
			}
		} else if choice == "2" {
			if !floatCalc() {
				break
			}
		} else if choice == "3" {
			break
		} else {
			fmt.Println("Invalid option, please enter 1, 2 or 3.")
		}
	}

	fmt.Println("Goodbye!")
}
