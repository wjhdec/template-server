package dbgorm

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

type Options struct {
	Dsn     string
	MaxOpen int
	MaxIdle int
	// MaxLifetime unit: second
	MaxLifetime int
}

func NewOptions(v *viper.Viper) *Options {
	o := new(Options)
	if err := v.UnmarshalKey("db", o); err != nil {
		panic(err)
	}
	return o
}

func New(cfg *Options) *DB {
	if db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{Logger: NewGormLog(true)}); err != nil {
		panic(err)
	} else {
		if origin, err := db.DB(); err != nil {
			panic(err)
		} else {
			origin.SetMaxOpenConns(cfg.MaxOpen)
			origin.SetMaxIdleConns(cfg.MaxIdle)
			origin.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
			return &DB{DB: db}
		}
	}
}
