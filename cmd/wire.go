//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package cmd

import (
	"github.com/google/wire"
	"github.com/wjhdec/template-server/internal/api/book"
	"github.com/wjhdec/template-server/pkg/dbgorm"
)

func newBookRouter() *book.Router {
	wire.Build(book.ProviderSet, newDB)
	return nil
}

func newDB() *dbgorm.DB {
	wire.Build(dbgorm.NewOptions, dbgorm.New, GetConfig)
	return nil
}
