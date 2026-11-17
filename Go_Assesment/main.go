package main

import "fmt"

func main() {
	inventory := NewInventory()

	items := []Item{
		{ID: 1, Name: "Keyboard", Quantity: 10, Price: 1200},
		{ID: 2, Name: "Mouse", Quantity: 15, Price: 500},
		{ID: 3, Name: "Monitor", Quantity: 5, Price: 8500},
	}

	for _, item := range items {
		if err := inventory.AddItem(item); err != nil {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("Initial Inventory:")
	inventory.DisplayItems()

	fmt.Println("\nUpdating Keyboard:")
	if err := inventory.UpdateItem(1, 20, 1500); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nRemoving Mouse:")
	if err := inventory.RemoveItem(2); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nCurrent Inventory:")
	inventory.DisplayItems()

	fmt.Println("\nInventory Summary:")

	summary := <-inventory.CalculateSummary()

	fmt.Println("Total Items:", summary.ItemCount)
	fmt.Printf("Total Inventory Value: %.2f\n", summary.TotalValue)

	fmt.Println("\nSearching for Monitor:")

	item, err := inventory.GetItem(3)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found: %+v\n", item)
	}
}
