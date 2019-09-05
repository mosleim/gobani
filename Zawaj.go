package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Zawaj :
type Zawaj struct {
	ID        uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
	Zauj      uint16     `gorm:"column:zauj" form:"zauj" json:"zauj"`
	Zaujah    uint16     `gorm:"column:zaujah" form:"zaujah" json:"zaujah"`
}

// TableName :
func (*Zawaj) TableName() string {
	return "zawajs"
}

// handler create
func initRoutersZawaj(r *gin.Engine, zawaj string) {
	route := r.Group("zawaj")
	route.GET("/", getZawajs)
	route.GET("/:id", getZawaj)
	route.POST("/", createZawaj)
	route.PUT("/:id", updateZawaj)
	route.DELETE("/:id", deleteZawaj)
}

func getZawajs(c *gin.Context) {
	var zawajs []Zawaj
	if err := g.Find(&zawajs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, zawajs)
	}
}

func getZawaj(c *gin.Context) {
	id := c.Params.ByName("id")
	var zawaj Zawaj
	if err := g.Where("id = ?", id).First(&zawaj).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, zawaj)
	}
}

func createZawaj(c *gin.Context) {
	var zawaj Zawaj
	c.BindJSON(&zawaj)
	g.Create(&zawaj)
	c.JSON(200, zawaj)
}

func updateZawaj(c *gin.Context) {
	var zawaj Zawaj
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&zawaj).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&zawaj)
	g.Save(&zawaj)
	c.JSON(200, zawaj)
}
func deleteZawaj(c *gin.Context) {
	id := c.Params.ByName("id")
	var zawaj Zawaj
	d := g.Where("id = ?", id).Delete(&zawaj)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
