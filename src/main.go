package main

import (
	"fmt"

	"github.com/daniel-munoz/shopping-list/src/model"
)

func main() {
	// Initialize categories with items
	groceries := model.Category{
		Name: "Groceries",
		Items: []model.Item{
			{Name: "bread"},
			{Name: "milk"},
			{Name: "eggs"},
			{Name: "butter"},
		},
	}
	fruits := model.Category{
		Name: "Fruits",
		Items: []model.Item{
			{Name: "Apple"},
			{Name: "Orange"},
			{Name: "Banana"},
		},
	}

	bakery := model.Category{
		Name: "Bakery",
		Items: []model.Item{
			{Name: "Bread"},
		},
	}

	dairy := model.Category{
		Name: "Dairy",
		Items: []model.Item{
			{Name: "Milk"},
		},
	}

	// Create a shopping list with the categories.
	shoppingList := model.ShoppingList{
		Categories: []model.Category{groceries, fruits, bakery, dairy},
	}
	// print the category objects:
	fmt.Printf("%v\n", shoppingList.String())
}
