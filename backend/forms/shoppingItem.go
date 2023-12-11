package forms

type CreateShoppingItemForm struct {
	Item   string `json:"item"  binding:"required"`
}