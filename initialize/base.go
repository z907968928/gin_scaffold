package initialize

import (
	"github.com/e421083458/gin_scaffold/core"
)

func InitBaseConf(confPath, confType string) error {
	// Parse Base Config
	if err := core.ParseConfig(confPath, confType, &core.ConfBase); err != nil {
		return err
	}

	return nil
}
