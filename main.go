package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shouva/dailyhelper"
)

func main() {
	// connect to db
	var config Config
	dailyhelper.ReadConfig(dailyhelper.GetCurrentPath(false)+"/config.json", &config)
	cDB := config.Database
	g, _ = open(cDB.Host, cDB.Port, cDB.DBName, cDB.User, cDB.Password)
	defer g.Close()
	migrate(g)

	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"selamat": "malam"})
	})
	loadrouter(r)
	r.Run()
}
