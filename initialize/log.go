package initialize

import (
	"github.com/e421083458/gin_scaffold/core"
	"github.com/e421083458/gin_scaffold/core/logger"
)

func InitLogConf() error {
	if err := logger.SetupDefaultLogWithConf(); err != nil {
		return err
	}
	logger.SetLayout(core.ConfBase.Log.Layout)
	return nil
}
