package controller

import (
	"database/sql"
	M "main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var shopping_item = new(M.Shopping_item)

// fetchItems retrieves items from the database
func FetchItems(db *sql.DB) ([]shopping_item, error) {
	rows, err := db.Query("SELECT id, item, bought, created FROM shopping_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []shopping_item
	for rows.Next() {
		var item shopping_item
		if err := rows.Scan(&item.ID, &item.Item, &item.Bought, &item.Created); err != nil {
			return nil, err
		}		
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func PostItem(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Param()
        var shopping_item shopping_item

        // Bind the incoming JSON to newItem struct
        if err := c.ShouldBindJSON(&shopping_item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Insert the new item into the database
        _, err := db.Exec("INSERT INTO shopping_items (item, bought) VALUES (?, ?)", shopping_item.Item, false)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert item into database"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Item added successfully"})
    }
}