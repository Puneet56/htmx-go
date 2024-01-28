package users

import "github.com/labstack/echo/v4"

var u = []User{
	{Id: 1, Name: "John Doe", Email: "test.com", Role: "Admin", Password: "123"},
	{Id: 2, Name: "Jane Doe", Email: "test.com", Role: "Project Manager", Password: "123"},
	{Id: 3, Name: "Jack Doe", Email: "test.com", Role: "Team Member", Password: "123"},
}

func Index(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return UsersIndex(u).Render(c.Request().Context(), c.Response().Writer)
}
