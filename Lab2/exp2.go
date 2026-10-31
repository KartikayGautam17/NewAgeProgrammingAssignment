package main

import (
	"fmt"

	"lib/lib"
)

func main() {
	word := "Hello, World"

	fmt.Println("Original:", word)
	fmt.Println("Reversed:", lib.Reverse(word))
	fmt.Println("Vowels:", lib.CountVowels(word))

	fmt.Println("5! =", lib.Factorial(5))
	fmt.Println("2^10 =", lib.Power(2, 10))
}
