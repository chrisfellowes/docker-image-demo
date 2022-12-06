package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("beginning test")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	path := os.Getenv("CUSTOM_PATH")
	message := os.Getenv("CUSTOM_MESSAGE")

	filePath := os.Getenv("CUSTOM_FILE_PATH")

	e.GET("/", func(c echo.Context) error {
		fmt.Println("got a /")
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		fmt.Println("got a /ping!")
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET(path, func(c echo.Context) error {
		fmt.Println("got a request at path")
		return c.JSON(http.StatusOK, message)
	})

	e.GET("/file", func(c echo.Context) error {
		fmt.Println("got a request at /file")
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, string(data))
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))

}
