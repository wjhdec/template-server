package dbgorm

import (
	"context"
	"strings"
	"time"

	"github.com/wjhdec/template-server/pkg/logger"
	gormLog "gorm.io/gorm/logger"
)

type gLog struct {
	SkipErrRecordNotFound bool
	log                   logger.Logger
}

func NewGormLog(skipRecordNotFound bool) gormLog.Interface {
	return &gLog{SkipErrRecordNotFound: skipRecordNotFound, log: logger.GetLogger()}
}

var _ gormLog.Interface = new(gLog)

func (g *gLog) LogMode(gormLog.LogLevel) gormLog.Interface {
	return g
}

func (g *gLog) Info(ctx context.Context, s string, i ...interface{}) {
	g.log.Infof(s, i...)
}

func (g *gLog) Warn(ctx context.Context, s string, i ...interface{}) {
	g.log.Warnf(s, i...)
}

func (g *gLog) Error(ctx context.Context, s string, i ...interface{}) {
	g.log.Errorf(s, i...)
}

type SqlInfo struct {
	Sql     string `json:"sql,omitempty"`
	Elapsed int64  `json:"elapsed,omitempty"`
	Error   error  `json:"error,omitempty"`
}

func (g *gLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	sqlInfo := &SqlInfo{Sql: sql, Elapsed: elapsed.Milliseconds()}
	if err != nil && !g.SkipErrRecordNotFound {
		sqlInfo.Error = err
		g.log.Errorf("json", sqlInfo)
		return
	}
	if strings.HasPrefix(strings.ToUpper(strings.TrimSpace(sql)), "SELECT") {
		g.log.Debugf("json", sqlInfo)
		return
	}

	if strings.HasPrefix(strings.ToUpper(strings.TrimSpace(sql)), "UPDATE") {
		g.log.Debugf("json", sqlInfo)
		return
	}
	g.log.Infof("json", sqlInfo)
}
