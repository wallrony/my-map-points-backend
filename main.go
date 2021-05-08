package main

import (
	"my_map_points/core/utils"
	"my_map_points/endpoint/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddlewares "github.com/labstack/echo/v4/middleware"
)

type HandlerRespose struct {
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}

func main() {
	echoInstance := echo.New()

	echoInstance.Use(echoMiddlewares.Logger())
	echoInstance.Use(echoMiddlewares.Recover())

	utils.LoadEnv()

	userHandler := handlers.NewUserHandler()

	echoInstance.GET("/users", userHandler.FindAll)

	echoInstance.Logger.Fatal(echoInstance.Start(":3333"))
}

func index(context echo.Context) error {
	response := &HandlerRespose{
		Message: "Hello world",
	}

	return context.JSON(http.StatusOK, response)
}
