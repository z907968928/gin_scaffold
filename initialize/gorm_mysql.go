package initialize

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/e421083458/gin_scaffold/core"
	LOG "github.com/e421083458/gin_scaffold/core/logger"
	"github.com/e421083458/gin_scaffold/lib"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type MysqlGormLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func InitDBPool(confPath, confType string) error {
	//普通的db方式

	if err := core.ParseConfig(confPath, confType, &core.ConfMysqlMap); err != nil {
		return err
	}

	if len(core.ConfMysqlMap.List) == 0 {
		LOG.PInfoF("%s%s\n", time.Now().Format(core.TimeFormat), " empty mysql config.")
	}

	if err := setDBPool(); err != nil {
		return err
	}

	return nil
}

func CloseDB() error {
	for _, dbPool := range core.DBMapPool {
		_ = dbPool.Close()
	}
	core.DBMapPool = make(map[string]*sql.DB)
	core.GORMMapPool = make(map[string]*gorm.DB)
	return nil
}

func setDBPool() error {
	core.DBMapPool = map[string]*sql.DB{}
	core.GORMMapPool = map[string]*gorm.DB{}
	for confName, DbConf := range core.ConfMysqlMap.List {
		connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&%s",
			DbConf.User,
			DbConf.Password,
			DbConf.Host,
			DbConf.Port,
			DbConf.Database,
			DbConf.Charset,
			DbConf.AddTo,
		)

		dbPool, err := sql.Open("mysql", connectInfo)
		if err != nil {
			return err
		}
		dbPool.SetMaxOpenConns(DbConf.MaxOpenConn)
		dbPool.SetMaxIdleConns(DbConf.MaxIdleConn)
		dbPool.SetConnMaxLifetime(time.Duration(DbConf.MaxConnLifeTime) * time.Second)
		err = dbPool.Ping()
		if err != nil {
			return err
		}

		//gorm连接方式
		dbGorm, err := gorm.Open(mysql.New(mysql.Config{Conn: dbPool}), &gorm.Config{
			Logger: &DefaultMysqlGormLogger,
		})
		if err != nil {
			return err
		}
		core.DBMapPool[confName] = dbPool
		core.GORMMapPool[confName] = dbGorm
	}

	//手动配置连接
	if dbpool, err := lib.GetDBPool("default"); err == nil {
		core.DBDefaultPool = dbpool
	}
	if dbpool, err := lib.GetGormPool("default"); err == nil {
		core.GORMDefaultPool = dbpool
	}
	return nil
}

//mysql 日志打印类型
var DefaultMysqlGormLogger = MysqlGormLogger{
	LogLevel:      logger.Info,
	SlowThreshold: 200 * time.Millisecond,
}

func (mgl *MysqlGormLogger) LogMode(logLevel logger.LogLevel) logger.Interface {
	mgl.LogLevel = logLevel
	return mgl
}

func (mgl *MysqlGormLogger) Info(ctx context.Context, message string, values ...interface{}) {
	trace := lib.GetTraceContext(ctx)
	params := make(map[string]interface{})
	params["message"] = message
	params["values"] = fmt.Sprint(values)
	lib.Log.TagInfo(trace, "_com_mysql_Info", params)
}

func (mgl *MysqlGormLogger) Warn(ctx context.Context, message string, values ...interface{}) {
	trace := lib.GetTraceContext(ctx)
	params := make(map[string]interface{})
	params["message"] = message
	params["values"] = fmt.Sprint(values)
	lib.Log.TagInfo(trace, "_com_mysql_Warn", params)
}

func (mgl *MysqlGormLogger) Error(ctx context.Context, message string, values ...interface{}) {
	trace := lib.GetTraceContext(ctx)
	params := make(map[string]interface{})
	params["message"] = message
	params["values"] = fmt.Sprint(values)
	lib.Log.TagInfo(trace, "_com_mysql_Error", params)
}

func (mgl *MysqlGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	trace := lib.GetTraceContext(ctx)

	if mgl.LogLevel <= logger.Silent {
		return
	}

	sqlStr, rows := fc()
	currentTime := begin.Format(core.TimeFormat)
	elapsed := time.Since(begin)
	msg := map[string]interface{}{
		"FileWithLineNum": utils.FileWithLineNum(),
		"sql":             sqlStr,
		"rows":            "-",
		"proc_time":       float64(elapsed.Milliseconds()),
		"current_time":    currentTime,
	}
	switch {
	case err != nil && mgl.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound)):
		msg["err"] = err
		if rows == -1 {
			lib.Log.TagInfo(trace, lib.DLTagMySqlFailed, msg)
		} else {
			msg["rows"] = rows
			lib.Log.TagInfo(trace, lib.DLTagMySqlFailed, msg)
		}
	case elapsed > mgl.SlowThreshold && mgl.SlowThreshold != 0 && mgl.LogLevel >= logger.Warn:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", mgl.SlowThreshold)
		msg["slowLog"] = slowLog
		if rows == -1 {
			lib.Log.TagInfo(trace, lib.DLTagMySqlSuccess, msg)
		} else {
			msg["rows"] = rows
			lib.Log.TagInfo(trace, lib.DLTagMySqlSuccess, msg)
		}
	case mgl.LogLevel == logger.Info:
		if rows == -1 {
			lib.Log.TagInfo(trace, lib.DLTagMySqlSuccess, msg)
		} else {
			msg["rows"] = rows
			lib.Log.TagInfo(trace, lib.DLTagMySqlSuccess, msg)
		}
	}
}
