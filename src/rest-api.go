package main

import (
	// "crypto"
	// "crypto/ed25519"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type EncryptRequest struct {
    EncryptionType string `json:"encryptionType"`
    PublicKey      string `json:"publicKey"`
    Data           string `json:"data"`
}

func createRoutes(router *gin.Engine) {
    router.POST("/encrypt", encrypt)
    router.POST("/decrypt", decrypt)
}

func encrypt(c *gin.Context) {
    request := EncryptRequest {}

    if err := c.BindJSON(&request); err != nil {
        fmt.Printf("An error occurred: %+v\n", err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    c.JSON(http.StatusOK, request)
}

func decrypt(c *gin.Context) {
    fmt.Println("Called decrypt")
}

func serveApi() {
    fmt.Println("Serving rest api in port 8082")

    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    config := cors.DefaultConfig()
    config.AllowWildcard = true
    config.AllowOrigins = []string{ "http://localhost:8080", "http://localhost:8080/*" }
    config.AllowMethods = []string{ "GET", "POST", "PUT", "DELETE" }
    router.Use(cors.New(config))

    createRoutes(router)

    err := router.Run("localhost:8082")
    
    if err != nil {
        fmt.Printf("%+v", err)
    }
}
