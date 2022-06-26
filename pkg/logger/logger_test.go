package logger

import (
	"testing"

	"github.com/labstack/echo/v4"
	gommonlog "github.com/labstack/gommon/log"
)

type Info struct {
	Sql string `json:"sql,omitempty"`
}

func TestLogInEcho(t *testing.T) {
	e := echo.New()
	SetLogger(e.Logger)
	e.Logger.SetLevel(gommonlog.DEBUG)
	log.Infof("json", &Info{Sql: "test"})
}
