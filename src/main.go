package main

import (
	"fmt"

	"github.com/daniel-munoz/shopping-list/src/model"
)

func main() {
	// Initialize categories with items
	fruits := model.NewCategory("Fruits", "Apple", "Orange", "Banana")

	bakery := model.NewCategory("Bakery", "Bread", "Bagel", "Donut")

	dairy := model.NewCategory("Dairy", "Milk", "Eggs", "Butter")

	// Select some items
	fruits.GetItem("Apple").Select()
	fruits.GetItem("Banana").Select()
	bakery.GetItem("Bread").Select()
	bakery.GetItem("Donut").Select()
	dairy.GetItem("Milk").Select()
	dairy.GetItem("Eggs").Select()

	// Buy some items
	fruits.GetItem("Apple").MarkAsBought()
	bakery.GetItem("Bread").MarkAsBought()
	bakery.GetItem("Donut").MarkAsBought()
	dairy.GetItem("Milk").MarkAsBought()


	// Create a shopping list with the categories.
	shoppingList := model.ShoppingList{
		Categories: []model.Category{fruits, bakery, dairy},
	}
	// print the category objects:
	fmt.Println("All items:")
	fmt.Printf("%v\n", shoppingList.String())
	// print the selected items:
	fmt.Println("Selected items:")
	fmt.Printf("%v\n", shoppingList.FilterSelected().String())
	// print the bought items:
	fmt.Println("Bought items:")
	fmt.Printf("%v\n", shoppingList.FilterBought().String())
}
