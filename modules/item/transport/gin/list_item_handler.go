package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"todo/common"
	"todo/modules/item/biz"
	"todo/modules/item/model"
	"todo/modules/item/storage"
)

func GetItems(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process()

		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var items []model.ToDoItem

		b := biz.NewItemListBiz(storage.NewSQLStore(db))
		items, err := b.GetItemListWithFilter(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(items, paging, filter))
	}

}
