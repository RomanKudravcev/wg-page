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

type ShoppingItemModel struct{}

//CREATE
func (m ShoppingItemModel) Create(form forms.CreateShoppingItemForm) (shoppingItemId int64, err error) {
    result, err := db.GetDB().Exec("INSERT INTO shopping_items (item, bought) VALUES (?, ?)", form.Item, false)
    if err != nil {
        return 0, err
    }

    shoppingItemId, err = result.LastInsertId()
    return shoppingItemId, err}

//GET ALL
func (m ShoppingItemModel) All() (shoppingItems []ShoppingItem, err error) {
	rows, err := db.GetDB().Query("SELECT id, item, bought, created FROM shopping_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item ShoppingItem
		if err := rows.Scan(&item.ID, &item.Item, &item.Bought, &item.Created); err != nil {
			return nil, err
		}		
		shoppingItems = append(shoppingItems, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return shoppingItems, err
}