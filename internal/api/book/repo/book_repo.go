package repo

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"github.com/wjhdec/template-server/internal/api/book/model"
	"github.com/wjhdec/template-server/pkg/dbgorm"
	"github.com/wjhdec/template-server/pkg/logger"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func (repo *BookRepo) Create(book *model.Book, ctx context.Context) (*model.Book, error) {
	if err := repo.db.WithContext(ctx).Create(book).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return book, nil
}

var repo *BookRepo
var once sync.Once

func New(db *dbgorm.DB) *BookRepo {
	once.Do(func() {
		repo = &BookRepo{db: db.DB}
		if err := repo.db.AutoMigrate(&model.Book{}); err != nil {
			logger.Fatalf("refresh talbe page error %+v", err)
		}
	})
	return repo
}
