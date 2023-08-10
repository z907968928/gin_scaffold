package lib

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/e421083458/gin_scaffold/core"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetDBPool(name string) (*sql.DB, error) {
	if dbPool, ok := core.DBMapPool[name]; ok {
		return dbPool, nil
	}
	return nil, errors.New("get pool error")
}

func GetGormPool(name string) (*gorm.DB, error) {
	if dbPool, ok := core.GORMMapPool[name]; ok {
		return dbPool, nil
	}
	return nil, errors.New("get pool error")
}

func DBPoolLogQuery(trace *TraceContext, sqlDb *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	startExecTime := time.Now()
	rows, err := sqlDb.Query(query, args...)
	endExecTime := time.Now()
	if err != nil {
		Log.TagError(trace, DLTagMySqlSuccess, map[string]interface{}{
			"sql":       query,
			"bind":      args,
			"proc_time": fmt.Sprintf("%f", endExecTime.Sub(startExecTime).Seconds()),
		})
	} else {
		Log.TagInfo(trace, DLTagMySqlSuccess, map[string]interface{}{
			"sql":       query,
			"bind":      args,
			"proc_time": fmt.Sprintf("%f", endExecTime.Sub(startExecTime).Seconds()),
		})
	}
	return rows, err
}
