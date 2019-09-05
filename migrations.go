package main

import "github.com/jinzhu/gorm"

func migrate(g *gorm.DB) {
	g.AutoMigrate(&Profile{})
	g.AutoMigrate(&Syakhsun{})
	g.AutoMigrate(&Zawaj{})
}
