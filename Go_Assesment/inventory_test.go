package main

import "testing"

func TestAddItem(t *testing.T) {
	inventory := NewInventory()

	item := Item{
		ID:       1,
		Name:     "Keyboard",
		Quantity: 10,
		Price:    1200,
	}

	err := inventory.AddItem(item)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(inventory.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(inventory.Items))
	}

	_, err = inventory.GetItem(1)

	if err != nil {
		t.Fatalf("item was not added correctly")
	}
}

func TestAddDuplicateItem(t *testing.T) {
	inventory := NewInventory()

	item := Item{
		ID:       1,
		Name:     "Mouse",
		Quantity: 5,
		Price:    500,
	}

	inventory.AddItem(item)

	err := inventory.AddItem(item)

	if err == nil {
		t.Fatal("expected duplicate item error")
	}
}

func TestUpdateItem(t *testing.T) {
	inventory := NewInventory()

	item := Item{
		ID:       1,
		Name:     "Monitor",
		Quantity: 5,
		Price:    8500,
	}

	inventory.AddItem(item)

	err := inventory.UpdateItem(1, 10, 9000)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updated, _ := inventory.GetItem(1)

	if updated.Quantity != 10 {
		t.Errorf("expected quantity 10, got %d", updated.Quantity)
	}

	if updated.Price != 9000 {
		t.Errorf("expected price 9000, got %.2f", updated.Price)
	}
}

func TestRemoveItem(t *testing.T) {
	inventory := NewInventory()

	item := Item{
		ID:       1,
		Name:     "Keyboard",
		Quantity: 10,
		Price:    1200,
	}

	inventory.AddItem(item)

	err := inventory.RemoveItem(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(inventory.Items) != 0 {
		t.Fatal("expected inventory to be empty")
	}
}

func TestInvalidItem(t *testing.T) {
	inventory := NewInventory()

	item := Item{
		ID:       0,
		Name:     "",
		Quantity: -5,
		Price:    -100,
	}

	err := inventory.AddItem(item)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if len(inventory.Items) != 0 {
		t.Fatal("invalid item should not be added")
	}
}

func TestCalculateSummary(t *testing.T) {
	inventory := NewInventory()

	inventory.AddItem(Item{
		ID:       1,
		Name:     "Keyboard",
		Quantity: 10,
		Price:    1200,
	})

	inventory.AddItem(Item{
		ID:       2,
		Name:     "Mouse",
		Quantity: 5,
		Price:    500,
	})

	summary := <-inventory.CalculateSummary()

	if summary.ItemCount != 2 {
		t.Errorf("expected 2 items, got %d", summary.ItemCount)
	}

	expectedValue := 14500.0

	if summary.TotalValue != expectedValue {
		t.Errorf(
			"expected total value %.2f, got %.2f",
			expectedValue,
			summary.TotalValue,
		)
	}
}
