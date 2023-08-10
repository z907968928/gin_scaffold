package initialize

import (
	"github.com/e421083458/gin_scaffold/core"
	"github.com/e421083458/gin_scaffold/core/logger"
	"time"
)

func InitProfile(confPath, confType string) error {
	confProfileMap := map[string]interface{}{
		"profile": &core.ConfProfile.ProfileConfig,
	}
	for model, confStruct := range confProfileMap {
		confPath, confType, err := core.GetConfPath(model)
		if err != nil {
			logger.PErrorF("%s init_%s: %s\n", time.Now().Format(core.TimeFormat), model, err.Error())
			return err
		}
		logger.PInfoF("init %s use conf file : %s", model, confPath)
		err = core.ParseConfig(confPath, confType, &confStruct)
		if err != nil {
			return err
		}
	}
	return nil
}
