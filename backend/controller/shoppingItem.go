package controller

import (
	"main/forms"
	M "main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ArticleController ...
type ShoppingItemController struct{}

var shoppingItem = new(M.ShoppingItem)

//Create ...
func (ctrl ShoppingItemController) Create(c *gin.Context) {
	var form forms.CreateShoppingItemForm


	id, err := M.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article created", "id": id})
}


// fetchItems retrieves items from the database
func (ctrl ShoppingItemController) All(c *gin.Context){
	results, err := M.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get articles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}