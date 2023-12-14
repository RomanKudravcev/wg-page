package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type CreateShoppingItemForm struct {
	Item string `form:"item" json:"item"  binding:"required"`
	Bought bool `form:"bought" json:"bought"`
}

type ShoppingItemForm struct{}

//Item
func (f ShoppingItemForm) Item(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the item name"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

//Bought
func (f ShoppingItemForm) Bought(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the item name"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

//CREATE
func (f ShoppingItemForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Item" {
				return f.Item(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

//UPDATE
func (f ShoppingItemForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Item" {
				return f.Item(err.Tag())
			}
			if err.Field() == "Bought" {
				return f.Bought(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}