package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gostack-labs/bytego"
)

var trans ut.Translator

func InitTrans(app *bytego.App, locale string) (err error) {
	zhTrans := zh.New()
	enTrans := en.New()

	uni := ut.New(zhTrans, zhTrans, enTrans)
	trans, ok := uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", locale)
	}
	v := validator.New()
	app.Validator(v.Struct)
	switch locale {
	case "zh":
		err = zh_trans.RegisterDefaultTranslations(v, trans)
		err = v.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
			return ut.Add("phone", "{0}必须是一个有效的手机号码！", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("phone", fe.Field())
			return t
		})
	case "en":
		err = en_trans.RegisterDefaultTranslations(v, trans)
	default:
		err = zh_trans.RegisterDefaultTranslations(v, trans)
	}
	v.RegisterValidation("phone", validPhone)
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		count := 2
		name := strings.SplitN(field.Tag.Get("json"), ",", count)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return
}

var validPhone validator.Func = func(fl validator.FieldLevel) bool {
	if phone, ok := fl.Field().Interface().(string); ok {
		result, err := regexp.MatchString(`^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`, phone)
		if err != nil {
			return false
		}
		return result
	}
	return true
}
