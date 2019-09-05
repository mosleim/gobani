package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Syakhsun :
type Syakhsun struct {
	ID        uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
	Name      string     `gorm:"column:name" form:"name" json:"name"`
	Aba       uint16     `gorm:"column:aba" form:"aba" json:"aba"`
	Ummi      uint16     `gorm:"column:ummi" form:"ummi" json:"ummi"`
	Gender    uint16     `gorm:"column:gender" form:"gender" json:"gender"`
	Birdday   *time.Time `gorm:"column:birdday" form:"birdday" json:"birdday"`
}

// TableName :
func (*Syakhsun) TableName() string {
	return "syakhsuns"
}

// handler create
func initRoutersSyakhsun(r *gin.Engine, syakhsun string) {
	route := r.Group("syakhsun")
	route.GET("/", getSyakhsuns)
	route.GET("/:id", getSyakhsun)
	route.POST("/", createSyakhsun)
	route.PUT("/:id", updateSyakhsun)
	route.DELETE("/:id", deleteSyakhsun)
}

func getSyakhsuns(c *gin.Context) {
	var syakhsuns []Syakhsun
	if err := g.Find(&syakhsuns).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, syakhsuns)
	}
}

func getSyakhsun(c *gin.Context) {
	id := c.Params.ByName("id")
	var syakhsun Syakhsun
	if err := g.Where("id = ?", id).First(&syakhsun).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, syakhsun)
	}
}

func createSyakhsun(c *gin.Context) {
	var syakhsun Syakhsun
	c.BindJSON(&syakhsun)
	g.Create(&syakhsun)
	c.JSON(200, syakhsun)
}

func updateSyakhsun(c *gin.Context) {
	var syakhsun Syakhsun
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&syakhsun).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&syakhsun)
	g.Save(&syakhsun)
	c.JSON(200, syakhsun)
}
func deleteSyakhsun(c *gin.Context) {
	id := c.Params.ByName("id")
	var syakhsun Syakhsun
	d := g.Where("id = ?", id).Delete(&syakhsun)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
