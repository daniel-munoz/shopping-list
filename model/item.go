package model

// Item represents an item in a shopping list.
type Item struct {
	Name     string `json:"name"`
	Selected bool   `json:"selected"`
	Bought   bool   `json:"bought"`
}

// NewItem creates a new item with the given name.
func NewItem(name string) *Item {
	return &Item{Name: name}
}

// GetName returns the name of the item.
func (i Item) GetName() string {
	return i.Name
}

// IsSelected returns true if the item is selected.
func (i Item) IsSelected() bool {
	return i.Selected
}

// IsBought returns true if the item is bought.
func (i Item) IsBought() bool {
	return i.Bought
}

// Select selects the item.
func (i *Item) Select() {
        i.Selected = true
}

// Unselect unselects the item.
func (i *Item) Unselect() {
	i.Selected = false
}

// MarkAsBought marks the item as bought.
func (i *Item) MarkAsBought() {
	i.Bought = true
}

// Unmark unmarks the item as bought.
func (i *Item) Unmark() {
	i.Bought = false
}

// String returns a string representation of the item.
func (i Item) String() string {
	return i.Name
}
