package initialize

import (
	"bytes"
	"github.com/e421083458/gin_scaffold/core"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

//初始化配置文件
func InitViperConf() error {
	f, err := os.Open(core.ConfEnvPath + "/")
	if err != nil {
		return err
	}
	fileList, err := f.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f0 := range fileList {
		if !f0.IsDir() {
			bts, err := ioutil.ReadFile(core.ConfEnvPath + "/" + f0.Name())
			if err != nil {
				return err
			}
			fileType := strings.Split(f0.Name(), ".")[1]
			v := viper.New()
			v.SetConfigType(fileType)
			if err := v.ReadConfig(bytes.NewBuffer(bts)); err != nil {
				return err
			}
			pathArr := strings.Split(f0.Name(), ".")
			if core.ViperConfMap == nil {
				core.ViperConfMap = make(map[string]*viper.Viper)
			}
			core.ViperConfMap[pathArr[0]] = v
		}
	}
	return nil
}
