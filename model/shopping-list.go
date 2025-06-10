package model

import (
	"fmt"
	"strings"
)

// ShoppingList represents a shopping list.
type ShoppingList struct {
	Categories []Category `json:"categories"`
}

// String returns a string representation of the shopping list.
func (sl ShoppingList) String() string {
	var b strings.Builder
	for _, category := range sl.Categories {
		fmt.Fprintf(&b, "%s\n", category)
	}
	return b.String()
}

// AddCategory adds a category to the shopping list.
func (sl *ShoppingList) AddCategory(category Category) {
	sl.Categories = append(sl.Categories, category)
}

// FilterSelected returns a new shopping list with only the selected items.
func (sl ShoppingList) FilterSelected() ShoppingList {
	var filtered ShoppingList
	for _, category := range sl.Categories {
		filteredCategory := category.FilterSelected()
		filtered.AddCategory(filteredCategory)
	}
	return filtered
}

// FilterBought returns a new shopping list with only the bought items.
func (sl ShoppingList) FilterBought() ShoppingList {
	var filtered ShoppingList
	for _, category := range sl.Categories {
		filteredCategory := category.FilterBought()
		filtered.AddCategory(filteredCategory)
	}
	return filtered
}
