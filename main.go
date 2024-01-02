package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Model struct {
	ID        int       `json:"id" gorm:"primarykey column:id;"`
	CreatedAt time.Time `json:"created_at" gorm:"column: created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column: updated_at;"`
}

type ToDoItemsModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Model
}

func (ToDoItemsModel) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Model
}

func (TodoItemCreation) TableName() string {
	return ToDoItemsModel{}.TableName()
}

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("")
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "create todo item success",
			"data":    data.ID,
		})
	}

}
func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data ToDoItemsModel

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data.ID = id
		if err := db.First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}

}
