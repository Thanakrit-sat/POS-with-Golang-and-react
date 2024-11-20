package main

import (
	h "ezpos/api/controller"
	"ezpos/api/controller/auth"
	jwtMiddleware "ezpos/api/middleware"
	"ezpos/api/model"
	"ezpos/api/postgresql"

	// "ezpos/api/model"
	_ "ezpos/api/model"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

}