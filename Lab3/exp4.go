package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

var reader = bufio.NewReader(os.Stdin)

func ask(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func (p *Person) ReadData() {
	p.Name = ask("Enter name: ")

	for {
		age, err := strconv.Atoi(ask("Enter age: "))
		if err != nil || age < 0 || age > 120 {
			fmt.Println("Please enter a valid age between 0 and 120.")
			continue
		}
		p.Age = age
		break
	}

	p.Job = ask("Enter job: ")

	for {
		salary, err := strconv.ParseFloat(ask("Enter salary: "), 64)
		if err != nil || salary < 0 {
			fmt.Println("Please enter a valid salary.")
			continue
		}
		p.Salary = salary
		break
	}
}

func (p Person) Display() {
	fmt.Println("-------------------------")
	fmt.Printf("Name   : %s\n", p.Name)
	fmt.Printf("Age    : %d\n", p.Age)
	fmt.Printf("Job    : %s\n", p.Job)
	fmt.Printf("Salary : %.2f\n", p.Salary)
	fmt.Println("-------------------------")
}

func main() {
	var first, second Person

	fmt.Println("Enter details for person 1")
	first.ReadData()
	fmt.Println()

	fmt.Println("Enter details for person 2")
	second.ReadData()
	fmt.Println()

	fmt.Println("Person 1:")
	first.Display()

	fmt.Println("Person 2:")
	second.Display()
}
