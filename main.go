package main

import (
	"htmxgo/handlers"
	"htmxgo/users"
	"htmxgo/views/home"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return home.Home().Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/projects", handlers.GetProjects)
	e.GET("/users", users.Index)
	e.Logger.Fatal(e.Start(":1323"))
}
