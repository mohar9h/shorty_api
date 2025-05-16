package validation

import (
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	fatranslations "github.com/go-playground/validator/v10/translations/fa"
)

var (
	Validator      *validator.Validate
	UniversalTrans *ut.UniversalTranslator
)

func InitValidator() error {
	Validator = validator.New()

	enLocale := en.New()
	faLocale := fa.New()
	UniversalTrans = ut.New(enLocale, enLocale, faLocale)

	for _, locale := range []string{"en", "fa"} {
		trans, _ := UniversalTrans.GetTranslator(locale)

		switch locale {
		case "fa":
			_ = fatranslations.RegisterDefaultTranslations(Validator, trans)

			// Custom Persian translations
			registerFaCustomTranslation("required", "{0} الزامی است", trans)
			registerFaCustomTranslation("email", "{0} باید یک ایمیل معتبر باشد", trans)
			registerFaCustomTranslation("min", "{0} باید حداقل {1} کاراکتر باشد", trans)
			registerFaCustomTranslation("max", "{0} باید حداکثر {1} کاراکتر باشد", trans)

		default:
			_ = entranslations.RegisterDefaultTranslations(Validator, trans)
		}
	}

	return nil
}

func GetTranslatorFromHeader(header string) ut.Translator {
	if strings.HasPrefix(header, "fa") {
		trans, _ := UniversalTrans.GetTranslator("fa")
		return trans
	}
	trans, _ := UniversalTrans.GetTranslator("en")
	return trans
}

func registerFaCustomTranslation(tag string, msg string, trans ut.Translator) {
	_ = Validator.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, msg, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field(), fe.Param())
		return t
	})
}
