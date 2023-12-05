package main

import (
    //"crypto"
    "fmt"

   "github.com/gin-gonic/gin"
)

func createRoutes(router *gin.Engine) {
    router.GET("/encrypt", encrypt)
    router.POST("/decrypt", encrypt)
}


func encrypt(c *gin.Context) {
    // pub_key string, data string
    fmt.Println("Called encrypt")
}

func decrypt(c *gin.Context) {
    fmt.Println("Called encrypt")
}

func serveApi() {
    fmt.Println("Serving rest api in port 8082")

    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    createRoutes(router)

    err := router.Run("localhost:8082")
    
    if err != nil {
        fmt.Printf("%+v", err)
    }
}
