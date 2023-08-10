package lib

import (
	"github.com/e421083458/gin_scaffold/core"
	"strings"
	"time"
)

//获取get配置信息
func GetStringConf(key string) string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return ""
	}
	v, ok := core.ViperConfMap[keys[0]]
	if !ok {
		return ""
	}
	confString := v.GetString(strings.Join(keys[1:len(keys)], "."))
	return confString
}

//获取get配置信息
func GetStringMapConf(key string) map[string]interface{} {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.GetStringMap(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func GetConf(key string) interface{} {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.Get(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func GetBoolConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.GetBool(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func GetFloat64Conf(key string) float64 {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.GetFloat64(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func GetIntConf(key string) int {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.GetInt(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func GetStringMapStringConf(key string) map[string]string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.GetStringMapString(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func GetStringSliceConf(key string) []string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}

	v := core.ViperConfMap[keys[0]]
	conf := v.GetStringSlice(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func GetTimeConf(key string) time.Time {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return time.Now()
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.GetTime(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取时间阶段长度
func GetDurationConf(key string) time.Duration {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.GetDuration(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//是否设置了key
func IsSetConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := core.ViperConfMap[keys[0]]
	conf := v.IsSet(strings.Join(keys[1:len(keys)], "."))
	return conf
}
