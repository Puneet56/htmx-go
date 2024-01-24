package users

import (
	"htmxgo/types"
	user "htmxgo/views/users"

	"github.com/labstack/echo/v4"
)

var users = []types.User{
	{Id: 1, Name: "John", Age: 20, City: "New York"},
	{Id: 2, Name: "Peter", Age: 25, City: "London"},
	{Id: 3, Name: "Mary", Age: 30, City: "Tokyo"},
	{Id: 4, Name: "Steve", Age: 35, City: "Singapore"},
}

func GetUsers(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return user.UsersPage(users).Render(c.Request().Context(), c.Response().Writer)
}
