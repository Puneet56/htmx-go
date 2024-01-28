package handlers

import (
	"htmxgo/types"
	"htmxgo/views"

	"github.com/labstack/echo/v4"
)

var p = make([]types.Project, 0)

func GetProjects(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return views.ProjectsPage(p).Render(c.Request().Context(), c.Response().Writer)
}
