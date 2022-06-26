package book

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wjhdec/template-server/internal/api/book/model"
	"github.com/wjhdec/template-server/internal/api/book/repo"
)

type Handler struct {
	repo *repo.BookRepo
}

func (h *Handler) Insert(c echo.Context) error {
	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		return errors.WithStack(err)
	}
	return c.JSON(http.StatusCreated, book)
}

func NewHandler(repo *repo.BookRepo) *Handler {
	return &Handler{repo: repo}
}
