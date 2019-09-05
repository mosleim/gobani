package main

import "github.com/jinzhu/gorm"

// Syakhsun :
type Syakhsun struct {
	ID   uint16
	Name string
	Aba  uint16
	Ummi uint16
}

// Zawaj :
type Zawaj struct {
	*gorm.Model
	Zauj   uint16
	Zaujah uint16
}

func main() {

}
