package core

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	_ "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 解析配置文件目录
// 配置文件必须放到一个文件夹中
// 如：config=conf/dev/base.json 	ConfEnvPath=conf/dev	ConfEnv=dev
// 如：config=conf/base.json		ConfEnvPath=conf		ConfEnv=conf
func ParseConfPath(config string) error {
	path := strings.Split(config, "/")
	prefix := strings.Join(path[:len(path)-1], "/")
	ConfEnvPath = prefix
	ConfEnv = path[len(path)-2]
	return nil
}

//获取配置环境名
func GetConfEnv() string {
	return ConfEnv
}

func GetConfPath(fileName string) (string, string, error) {
	var (
		filePath string
	)
	for _, fileType := range ConfEnvFileType {
		confFleName := fmt.Sprintf("%s.%s", fileName, fileType)
		filePathNames, _ := filepath.Glob(filepath.Join(ConfEnvPath, confFleName))
		if len(filePathNames) >= 1 {
			filePath = fmt.Sprintf("%s/%s.%s", ConfEnvPath, fileName, fileType)
			return filePath, fileType, nil

		}
	}
	return "", "", errors.New("conf file not found")
}

func ParseConfig(confPath, confType string, conf interface{}) error {
	file, err := os.Open(confPath)
	if err != nil {
		return fmt.Errorf("open config %v fail, %v", confPath, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("read config fail, %v", err)
	}
	v := viper.New()
	a := strings.Split(confPath, "/")
	fn := strings.Split(a[len(a)-1], ".")[0]
	v.AddConfigPath(ConfEnvPath)
	v.SetConfigName(fn)
	v.SetConfigType(confType)
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read config fail, %v", err)
	}

	//if err := v.ReadConfig(bytes.NewBuffer(data)); err != nil {
	//	return fmt.Errorf("read config fail, %v", err)
	//}

	if err := v.Unmarshal(conf); err != nil {
		return fmt.Errorf("parse config fail, config:%v, err:%v", string(data), err)
	}
	return nil
}
