package main

import "fmt"

type Student struct {
	Name  string
	Roll  int
	Marks float64
}

func addBonus(score *int, bonus int) {
	*score = *score + bonus
}

func main() {
	fmt.Println("--- Address and value ---")
	num := 42
	ptr := &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num (&num):", &num)
	fmt.Println("Pointer ptr holds:", ptr)
	fmt.Println("Value through pointer (*ptr):", *ptr)

	*ptr = 100
	fmt.Println("After *ptr = 100, num is:", num)

	fmt.Println()
	fmt.Println("--- Pass by reference ---")
	score := 70
	fmt.Println("Before:", score)
	addBonus(&score, 15)
	fmt.Println("After:", score)

	fmt.Println()
	fmt.Println("--- Struct with new() ---")
	s := new(Student)
	fmt.Println("Zero value:", *s)

	s.Name = "Ananya"
	s.Roll = 21
	s.Marks = 88.5
	fmt.Println("After setting fields:", *s)

	s.Marks += 5
	fmt.Printf("Updated marks for %s (roll %d): %.1f\n", s.Name, s.Roll, s.Marks)
}
