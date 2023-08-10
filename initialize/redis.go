package initialize

import (
	"github.com/e421083458/gin_scaffold/core"
)

func InitRedisConf(confPath, confType string) error {
	err := core.ParseConfig(confPath, confType, &core.ConfRedisMap)
	if err != nil {
		return err
	}
	return nil
}
