package main

import "github.com/jinzhu/gorm"
//import _ "github.com/go-sql-driver/mysql"
import _ "github.com/jinzhu/gorm/dialects/mysql"
import "time"
// Syakhsun :
type Syakhsun struct {
	*gorm.Model
	Name string `gorm:"name" json:"name"`
	Aba  uint16 `gorm:"aba" json:"aba"`
	Ummi uint16 `gorm:"ummi" json:"ummi"`
	Gender uint16 `gorm:"gender" json:"gender"`
	Birdday *time.Time
}

// Zawaj :
type Zawaj struct {
	*gorm.Model
	Zauj   uint16 `gorm:"zauj" json:"zauj"`
	Zaujah uint16 `gorm:"zaujah" json:"zaujah"`
}

// Profile :
type Profile struct {
	*gorm.Model
	Phone string 
	Address string
	Coordinate string 
	
}


func main() {
  db, _ := gorm.Open("mysql", "bani:rawi#123@/bani?charset=utf8&parseTime=True&loc=Local")
  db.AutoMigrate(&Syakhsun{}, &Zawaj{}, &Profile{})
  defer db.Close()
}
