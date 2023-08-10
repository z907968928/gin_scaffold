package core

import (
	"database/sql"
	"github.com/e421083458/gin_scaffold/core/conf"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net"
	"time"
)

var (
	TimeLocation *time.Location
	TimeFormat   = "2006-01-02 15:04:05"
	DateFormat   = "2006-01-02"
	LocalIP      = net.ParseIP("127.0.0.1")
)

var (
	ConfEnvPath     string //配置文件夹
	ConfEnv         string //配置环境名 比如：dev prod test
	ConfEnvFileType = []string{
		"yaml", "yml", "toml", "json",
	} //配置文件优先级
)

var (
	ConfBase     *conf.BaseConf
	ConfRedisMap *conf.RedisConf
	ConfMysqlMap *conf.MysqlConf
	ViperConfMap map[string]*viper.Viper
)

var (
	DBMapPool       map[string]*sql.DB
	GORMMapPool     map[string]*gorm.DB
	DBDefaultPool   *sql.DB
	GORMDefaultPool *gorm.DB
)

var (
	ConfProfile struct {
		ProfileConfig *conf.ProfileConfig
	}
)
