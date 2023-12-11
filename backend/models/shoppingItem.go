package models

import (
	"main/db"
	"main/forms"
)

type ShoppingItem struct {
	ID     int    `json:"id"`
	Item   string `json:"item"`
	Bought bool   `json:"bought"`
	Created string `json:"created"`
}

//Create ...
func Create(form forms.CreateShoppingItemForm) (shoppingItemId int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO shopping_items (item, bought) VALUES (?, ?)", form.Item, false).Scan(&shoppingItemId)
	return shoppingItemId, err
}
