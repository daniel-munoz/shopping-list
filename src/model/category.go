package model

import (
	"fmt"
	"strings"
)

// Category represents a category of items.
type Category struct {
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

// String returns a string representation of the category.
func (c Category) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s: {\n", c.Name)
	for _, item := range c.Items {
		fmt.Fprintf(&b, "\t[ ] %s\n", item.Name)
	}
	fmt.Fprintf(&b, "}")
	return b.String()
}

// IsEmpty checks whether the category has no elements.
func (c Category) IsEmpty() bool {
	return len(c.Items) == 0
}
