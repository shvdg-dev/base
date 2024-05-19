package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

type Translator struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

func NewTranslator() *Translator {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	localizer := i18n.NewLocalizer(bundle, English)
	return &Translator{bundle: bundle, localizer: localizer}
}

func (t *Translator) AddTranslations(translationPath ...string) {
	for _, path := range translationPath {
		_, err := t.bundle.LoadMessageFile(path)
		if err != nil {
			log.Fatalf("could not load message file: %v", err)
		}
	}
}

func (t *Translator) Translate(key string) string {
	localized, err := t.localizer.Localize(&i18n.LocalizeConfig{MessageID: key})
	if err != nil {
		log.Printf("could not get text for '%s'", key)
		return "☒"
	}
	return localized
}
