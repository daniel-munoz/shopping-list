package model

import (
	"fmt"
	"strings"
)

// Category represents a category of items.
type Category struct {
	Name  string           `json:"name"`
	Items map[string]*Item `json:"items"`
}

// NewCategory creates a new category with the given name.
func NewCategory(name string, items ...string) Category {
	c := Category{
		Name:  name,
		Items: make(map[string]*Item),
	}
	for _, item := range items {
		c.AddItem(NewItem(item))
	}
	return c
}

// GetItem returns the item with the given name.
func (c Category) GetItem(name string) *Item {
	return c.Items[name]
}

// Len returns the number of items in the category.
func (c Category) Len() int {
	return len(c.Items)
}

// AddItem adds an item to the category.
func (c *Category) AddItem(item *Item) {
	c.Items[item.Name] = item
}

// String returns a string representation of the category.
func (c Category) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s: {\n", c.Name)
	for _, item := range c.Items {
		fmt.Fprintf(&b, "\t[")
		if item.IsBought() {
			fmt.Fprintf(&b, "X] %s\n", item.Name)
		} else {
			fmt.Fprintf(&b, " ] %s\n", item.Name)
		}
	}
	fmt.Fprintf(&b, "}")
	return b.String()
}

// IsEmpty checks whether the category has no elements.
func (c Category) IsEmpty() bool {
	return len(c.Items) == 0
}

// FilterBought returns a new category containing only the items that have been bought.
func (c Category) FilterBought() Category {
	var bought Category = NewCategory(c.Name)
	for _, item := range c.Items {
		if item.Bought {
			bought.AddItem(item)
		}
	}
	return bought
}

// FilterNotBought returns a new category containing only the items that have not been bought.
func (c Category) FilterNotBought() Category {
	var notBought Category = NewCategory(c.Name)
	for _, item := range c.Items {
		if !item.Bought {
			notBought.AddItem(item)
		}
	}
	return notBought
}

// FilterSelected returns a new category containing only the items that have been selected.
func (c Category) FilterSelected() Category {
	var selected Category = NewCategory(c.Name)
	for _, item := range c.Items {
		if item.Selected {
			selected.AddItem(item)
		}
	}
	return selected
}
