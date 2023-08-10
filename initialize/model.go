package initialize

import (
	"flag"
	"github.com/e421083458/gin_scaffold/core"
	"github.com/e421083458/gin_scaffold/core/logger"
	"github.com/e421083458/gin_scaffold/lib"
	"os"
	"time"
)

var (
	InitModuleMap = map[string]func(confPath, confType string) error{
		"base":      InitBaseConf,
		"mysql_map": InitDBPool,
		"redis_map": InitRedisConf,
		"profile":   InitProfile,
	}
)

//模块初始化
func InitModule(configPath string, modules []string) error {
	conf := flag.String("config", configPath, "input config file like ./conf/dev/")
	flag.Parse()
	if *conf == "" {
		flag.Usage()
		os.Exit(1)
	}

	logger.Println("------------------------------------------------------------------------")
	logger.PInfoF("config=%s\n", *conf)
	logger.PInfoF("%s\n", "start loading resources.")

	// 设置ip信息，优先设置便于日志打印
	ips := lib.GetLocalIPs()
	if len(ips) > 0 {
		core.LocalIP = ips[0]
	}

	// 解析配置文件目录
	if err := core.ParseConfPath(*conf); err != nil {
		return err
	}

	for _, model := range modules {
		if _, ok := InitModuleMap[model]; ok {
			logger.PInfoF("start init %s", model)
			confPath, confType, err := core.GetConfPath(model)
			logger.PInfoF("init %s use conf file : %s", model, confPath)
			if err != nil {
				logger.PErrorF("%s init_%s: %s\n", time.Now().Format(core.TimeFormat), model, err.Error())
				return err
			}
			if err := InitModuleMap[model](confPath, confType); err != nil {
				logger.PErrorF("%s init_%s: %s\n", time.Now().Format(core.TimeFormat), model, err.Error())
				return err
			}
		}
	}

	//初始化配置文件
	logger.PInfo("init viper")
	if err := InitViperConf(); err != nil {
		return err
	}

	// 初始化日志配置
	logger.PInfo("init logger")
	if err := InitLogConf(); err != nil {
		return err
	}

	// 设置时区
	if location, err := time.LoadLocation(core.ConfBase.Base.TimeLocation); err != nil {
		return err
	} else {
		core.TimeLocation = location
	}
	logger.PInfoF("%s\n", " success loading resources.")
	logger.Println("------------------------------------------------------------------------")
	return nil
}

//公共销毁函数
func Destroy() {
	logger.Println("------------------------------------------------------------------------")
	logger.PInfoF("%s\n", " start destroy resources.")
	if err := CloseDB(); err != nil {
		logger.PErrorF("%s\n", "dstroy db error.")
	}
	logger.Close()
	logger.PInfoF("%s\n", " success destroy resources.")
}
