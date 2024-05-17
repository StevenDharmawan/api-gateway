package main

import (
	"api-gateway/handler"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	company := os.Getenv("COMPANY_SERVICE")
	product := os.Getenv("PRODUCT_SERVICE")
	sales := os.Getenv("SALES_SERVICE")
	r := gin.Default()

	r.Use(middleware.JWTMiddleware())

	r.Any("/proxy/*proxyPath", handler.ReverseProxy(company))
	r.Any("/proxy/*proxyPath", handler.ReverseProxy(product))
	r.Any("/proxy/*proxyPath", handler.ReverseProxy(sales))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
