package main

import (
	"errors"
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type Inventory struct {
	Items []Item
}

type Summary struct {
	ItemCount  int
	TotalValue float64
}

func NewInventory() *Inventory {
	return &Inventory{
		Items: make([]Item, 0),
	}
}

func validateItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("item ID must be positive")
	}

	if item.Name == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	if item.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	return nil
}

func (inv *Inventory) AddItem(item Item) error {
	if err := validateItem(item); err != nil {
		return err
	}

	for _, existing := range inv.Items {
		if existing.ID == item.ID {
			return errors.New("item already exists")
		}
	}

	inv.Items = append(inv.Items, item)
	return nil
}

func (inv *Inventory) UpdateItem(id int, quantity int, price float64) error {
	if id <= 0 {
		return errors.New("invalid item ID")
	}

	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	if price <= 0 {
		return errors.New("price must be greater than zero")
	}

	for i := range inv.Items {
		if inv.Items[i].ID == id {
			inv.Items[i].Quantity = quantity
			inv.Items[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

func (inv *Inventory) RemoveItem(id int) error {
	if id <= 0 {
		return errors.New("invalid item ID")
	}

	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items = append(
				inv.Items[:i],
				inv.Items[i+1:]...,
			)
			return nil
		}
	}

	return errors.New("item not found")
}

func (inv *Inventory) GetItem(id int) (Item, error) {
	for _, item := range inv.Items {
		if item.ID == id {
			return item, nil
		}
	}

	return Item{}, errors.New("item not found")
}

func (inv *Inventory) ListItems() []Item {
	items := make([]Item, len(inv.Items))
	copy(items, inv.Items)
	return items
}

func (inv *Inventory) CalculateSummary() <-chan Summary {
	result := make(chan Summary)

	go func() {
		defer close(result)

		summary := Summary{
			ItemCount: len(inv.Items),
		}

		for _, item := range inv.Items {
			summary.TotalValue += float64(item.Quantity) * item.Price
		}

		result <- summary
	}()

	return result
}

func (inv *Inventory) DisplayItems() {
	for _, item := range inv.Items {
		fmt.Printf(
			"ID: %d | Name: %s | Quantity: %d | Price: %.2f\n",
			item.ID,
			item.Name,
			item.Quantity,
			item.Price,
		)
	}
}
