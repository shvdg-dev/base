package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

type I18n struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

func NewConfig() *I18n {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, _ = bundle.LoadMessageFile("resources/texts/en.toml")
	localizer := i18n.NewLocalizer(bundle, English)
	return &I18n{bundle: bundle, localizer: localizer}
}

func (i *I18n) GetText(key string) string {
	localized, err := i.localizer.Localize(&i18n.LocalizeConfig{MessageID: key})
	if err != nil {
		log.Printf("could not get text for '%s'", key)
		return "☒"
	}
	return localized
}
