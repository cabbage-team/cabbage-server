package i18n

import (
	"cabbage-server/model"

	"github.com/jingyuexing/go-utils"
	"github.com/jingyuexing/i18n"
)

var I18N *i18n.I18n

func init(){
	I18N = i18n.CreateI18n(&i18n.I18n{
		Local: "zh",
		Message: map[string]any{
			"zh":utils.LoadConfig[model.Translate]("locale/zh.json"),
			"en":utils.LoadConfig[model.Translate]("locale/en.json"),
		},
		FallbackLocale: "en",
	})

}