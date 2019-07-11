package rsserver

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
	"net/http"
)

type ItemController struct {
	Table *DB
}

func FetchItems(items *sql.Rows, ret *[]*gofeed.Item) (err error) {
	for items.Next() {
		var item gofeed.Item
		if err = items.Scan(&item.GUID, &item.Title, &item.Link, &item.Description, &item.Published); err != nil {
			return err
		}
		*ret = append(*ret, &item)
	}

	return nil
}

func (controller *ItemController) GetItems(ctx *gin.Context) {
	var err error
	var items []*gofeed.Item

	itemRows, err := controller.Table.conn.Query(`SELECT guid, title, link, description, pub_date FROM Item`)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err := FetchItems(itemRows, &items); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, items)
		}
	}
}
