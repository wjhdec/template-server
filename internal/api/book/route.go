package book

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/wjhdec/template-server/internal/api/book/repo"
)

type Router struct {
	handler *Handler
}

func (r *Router) Route(group *echo.Group) {
	bookGroup := group.Group("books")
	bookGroup.POST("", r.handler.Insert)
}

func New(handler *Handler) *Router {
	return &Router{handler: handler}
}

var ProviderSet = wire.NewSet(New, NewHandler, repo.New)
