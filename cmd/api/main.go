package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Printf("Error loading .env file: %v", err)
    }

    router := gin.Default()

    log.Fatal(router.Run(":8080"))
}
