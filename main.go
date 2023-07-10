package main

import (
	"github.com/gin-gonic/gin"

	"keuangan/connection"
	"keuangan/routers"
)

func main() {
	r := gin.Default()
	db := connection.Connection()
	redis := connection.Redis()

	eng := &routers.Routes{
		Db:    db,
		R:     r,
		Redis: redis,
	}

	eng.Routers()

	r.Run(":8084")
}
