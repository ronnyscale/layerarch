package main

import (
    "github.com/gin-gonic/gin"
    "github.com/ronnyscale/layerarch/internal/handlers"
    "github.com/ronnyscale/layerarch/pkg/database"
    "github.com/ronnyscale/layerarch/config"
)

func main() {
    err := db.InitDB(config.DSN)
    if err != nil {
        panic("failed to connect database")
    }

    r := gin.Default()
    r.POST("/users", handlers.CreateUser)
    r.Run(":8080")
}