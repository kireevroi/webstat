package main

import (
	"github.com/gin-gonic/gin"
	
	"github.com/kireevroi/webstat/internal/vdb"
	"github.com/kireevroi/webstat/internal/endpoints"
	"github.com/kireevroi/webstat/internal/statistics"
)

func main() {
	d := &vdb.DataBase{}
	go d.Init("list.txt")

	pingstat := &statistics.StatMap{}
	pingstat.Init()
	maxstat := &statistics.Stats{}
	minstat := &statistics.Stats{}

	router := gin.Default()
	router.Use(endpoints.ApiMiddleware("12345"))
	router.GET("/api/ping", endpoints.WebsiteTimeHandler(d, pingstat))
	router.GET("/api/max", endpoints.MaxHandler(d, maxstat))
	router.GET("/api/min", endpoints.MinHandler(d, minstat))
	router.Run(":8080")
}