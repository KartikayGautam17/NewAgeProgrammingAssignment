package main

import (
	"fmt"
	"sort"
)

func addStudent(names []string, name string) []string {
	return append(names, name)
}

func removeAt(names []string, index int) []string {
	if index < 0 || index >= len(names) {
		fmt.Println("Index", index, "is out of range, nothing removed")
		return names
	}
	return append(names[:index], names[index+1:]...)
}

func updateAt(names []string, index int, name string) {
	if index < 0 || index >= len(names) {
		fmt.Println("Index", index, "is out of range, nothing updated")
		return
	}
	names[index] = name
}

func printMarks(marks map[string]int) {
	keys := make([]string, 0, len(marks))
	for k := range marks {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Print("map[")
	for i, k := range keys {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%s:%d", k, marks[k])
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("--- Slice operations ---")

	students := []string{"Aarav", "Diya", "Kabir"}
	fmt.Println("Initial:", students)

	students = addStudent(students, "Meera")
	fmt.Println("After add Meera:", students)

	students = addStudent(students, "Rohan")
	fmt.Println("After add Rohan:", students)

	students = removeAt(students, 1)
	fmt.Println("After remove index 1:", students)

	updateAt(students, 0, "Arjun")
	fmt.Println("After update index 0 to Arjun:", students)

	students = removeAt(students, 10)
	fmt.Println("After invalid remove:", students)

	fmt.Println("Length:", len(students), "Capacity:", cap(students))

	fmt.Println()
	fmt.Println("--- Map operations ---")

	marks := map[string]int{
		"Maths":   88,
		"Physics": 76,
	}
	fmt.Print("Initial: ")
	printMarks(marks)

	marks["Chemistry"] = 91
	fmt.Print("After insert Chemistry: ")
	printMarks(marks)

	marks["English"] = 69
	fmt.Print("After insert English: ")
	printMarks(marks)

	marks["Physics"] = 80
	fmt.Print("After update Physics: ")
	printMarks(marks)

	delete(marks, "English")
	fmt.Print("After delete English: ")
	printMarks(marks)

	subjects := []string{"Maths", "English"}
	for _, s := range subjects {
		value, found := marks[s]
		if found {
			fmt.Printf("Lookup %s: found, marks = %d\n", s, value)
		} else {
			fmt.Printf("Lookup %s: not found\n", s)
		}
	}

	fmt.Println("Total subjects:", len(marks))
}
