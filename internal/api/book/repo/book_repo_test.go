package repo

import (
	"testing"

	"github.com/wjhdec/template-server/pkg/dbgorm"
	"github.com/wjhdec/template-server/pkg/logger"
)

func TestConnectBook(t *testing.T) {
	dsn := "host=localhost user=book password=book dbname=book port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db := dbgorm.New(&dbgorm.Options{Dsn: dsn})
	bookRepo := New(db)
	logger.Debugf(bookRepo.db.Name())
}
