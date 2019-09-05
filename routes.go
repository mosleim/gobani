package main

import "github.com/gin-gonic/gin"

func loadrouter(rgin *gin.Engine) {
	initRoutersProfile(rgin, "profile")
	initRoutersSyakhsun(rgin, "syakhsun")
	initRoutersZawaj(rgin, "zawaj")
}
