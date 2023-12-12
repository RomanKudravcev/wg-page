package controller

import (
	"main/forms"
	M "main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ArticleController ...
type ShoppingItemController struct{}

var shoppingItemForm = new(forms.ShoppingItemForm)
var shoppingItemModel = new(M.ShoppingItemModel)

//CREATE
func (ctrl ShoppingItemController) Create(c *gin.Context) {
	var form forms.CreateShoppingItemForm


	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := shoppingItemForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := shoppingItemModel.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "ShoppingItem could not be created","error:":err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ShoppingItem created", "id": id})
}


// fetchItems retrieves items from the database
func (ctrl ShoppingItemController) All(c *gin.Context){
	results, err := shoppingItemModel.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get shoppingItems"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": results})
}