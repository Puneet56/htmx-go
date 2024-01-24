package main

import (
	users "htmxgo/handlers"
	"htmxgo/views/home"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return home.Home().Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/users", users.GetUsers)
	e.Logger.Fatal(e.Start(":1323"))
}
