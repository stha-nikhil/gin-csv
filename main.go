package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stha-nikhil/gin-csv/api"
)

func main() {
	router := gin.Default()
	router.GET("/csv-create", api.CreateCSV)
	router.POST("/csv-fetch", api.GetCSV)
	router.Run()
}
