package api

import "github.com/labstack/echo/v4"

type Router interface {
	Route(g *echo.Group)
}
