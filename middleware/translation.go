package middleware

import (
	"github.com/e421083458/gin_scaffold/core"
	"github.com/e421083458/gin_scaffold/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type validateFuncMap struct {
	Tran string
	Func func(fl validator.FieldLevel) bool
}

var (
	validateFunc = map[string]validateFuncMap{
		"isIntIds": {
			"输入格式错误(1,2)",
			isIntIds,
		},
		"isDate": {
			"格式错误(YYYY-mm-dd HH:mm:ss)",
			isDate,
		},
		"isDateTime": {
			"格式错误(YYYY-mm-dd HH:mm:ss)",
			isDateTime,
		},
	}
)

//设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		//设置支持语言
		en := en.New()
		zh := zh.New()

		//设置国际化翻译器
		uni := ut.New(zh, zh, en)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			_ = zh_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			_ = registerValidation(val, trans)

			break
		}

		c.Set(utils.TranslatorKey, trans)
		c.Set(utils.ValidatorKey, val)
		c.Next()
	}
}

func registerValidation(val *validator.Validate, trans ut.Translator) error {
	//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
	for tag, regFunc := range validateFunc {
		// 方法注册
		val.RegisterValidation(tag, regFunc.Func)

		// 错误输出
		val.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
			return ut.Add(tag, "{0}"+regFunc.Tran, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field())
			return t
		})
	}
	return nil
}

// ids  (111,222)
func isIntIds(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return true
	}
	preg := `^[1-9][0-9]*$`
	values := strings.Split(value, ",")
	for i := 0; i < len(values); i++ {
		reg := regexp.MustCompile(preg)
		check := reg.MatchString(values[i])
		if !check {
			return false
		}
	}
	return true
}

func isDate(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return true
	}
	_, err := time.ParseInLocation(core.DateFormat, value, core.TimeLocation)
	if err != nil {
		return false
	}
	return true
}

func isDateTime(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return true
	}
	_, err := time.ParseInLocation(core.TimeFormat, value, core.TimeLocation)
	if err != nil {
		return false
	}
	return true
}
