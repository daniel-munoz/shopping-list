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
