package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Profile :
type Profile struct {
	ID         uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	CreatedAt  *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
	Phone      string     `gorm:"column:phone" form:"phone" json:"phone"`
	Address    string     `gorm:"column:address" form:"address" json:"address"`
	Coordinate string     `gorm:"column:coordinate" form:"coordinate" json:"coordinate"`
}

// TableName :
func (*Profile) TableName() string {
	return "profiles"
}

// handler create
func initRoutersProfile(r *gin.Engine, profile string) {
	route := r.Group("profile")
	route.GET("/", getProfiles)
	route.GET("/:id", getProfile)
	route.POST("/", createProfile)
	route.PUT("/:id", updateProfile)
	route.DELETE("/:id", deleteProfile)
}

func getProfiles(c *gin.Context) {
	var profiles []Profile
	if err := g.Find(&profiles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, profiles)
	}
}

func getProfile(c *gin.Context) {
	id := c.Params.ByName("id")
	var profile Profile
	if err := g.Where("id = ?", id).First(&profile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, profile)
	}
}

func createProfile(c *gin.Context) {
	var profile Profile
	c.BindJSON(&profile)
	g.Create(&profile)
	c.JSON(200, profile)
}

func updateProfile(c *gin.Context) {
	var profile Profile
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&profile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&profile)
	g.Save(&profile)
	c.JSON(200, profile)
}
func deleteProfile(c *gin.Context) {
	id := c.Params.ByName("id")
	var profile Profile
	d := g.Where("id = ?", id).Delete(&profile)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
