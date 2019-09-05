package main

import "github.com/jinzhu/gorm"

// Syakhsun :
type Syakhsun struct {
	*gorm.Model
	Name string `gorm:"name" json:"name"`
	Aba  uint16 `gorm:"aba" json:"aba"`
	Ummi uint16 `gorm:"ummi" json:"ummi"`
}

// Zawaj :
type Zawaj struct {
	*gorm.Model
	Zauj   uint16 `gorm:"zauj" json:"zauj"`
	Zaujah uint16 `gorm:"zaujah" json:"zaujah"`
}

func main() {

}
