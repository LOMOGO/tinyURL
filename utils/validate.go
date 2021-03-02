package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

func Valid(err error) (bool, map[string]string) {
	msg := make(map[string]string)
	if err != nil {
		uni := ut.New(zh.New())
		trans, _ := uni.GetTranslator("zh")
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			_ = zh2.RegisterDefaultTranslations(v, trans)
		}
		validErrs := err.(validator.ValidationErrors)

		for key, value := range validErrs.Translate(trans) {
			msg[key] = value
		}
		return false, msg
	}
	return true, nil
}
