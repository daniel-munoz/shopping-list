package main

import (
	"fmt"

	"github.com/daniel-munoz/shopping-list/src/model"
)

func main() {
	// Initialize categories with items
	groceries := model.NewCategory("Groceries", "Bread", "Milk", "Eggs", "Butter")

	fruits := model.NewCategory("Fruits", "Apple", "Orange", "Banana")

	bakery := model.NewCategory("Bakery", "Bread")

	dairy := model.NewCategory("Dairy", "Milk")

	// Create a shopping list with the categories.
	shoppingList := model.ShoppingList{
		Categories: []model.Category{groceries, fruits, bakery, dairy},
	}
	// print the category objects:
	fmt.Printf("%v\n", shoppingList.String())
}
