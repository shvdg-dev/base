package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

type Localizer struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

func NewLocalizer() *Localizer {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	localizer := i18n.NewLocalizer(bundle, "en")
	return &Localizer{bundle: bundle, localizer: localizer}
}

func (l *Localizer) AddLocalization(messageFilePath ...string) {
	for _, path := range messageFilePath {
		_, err := l.bundle.LoadMessageFile(path)
		if err != nil {
			log.Fatalf("could not load message file: %v", err)
		}
	}
}

func (l *Localizer) Localize(key string) string {
	localized, err := l.localizer.Localize(&i18n.LocalizeConfig{MessageID: key})
	if err != nil {
		log.Printf("could not get text for '%s'", key)
		return "☒"
	}
	return localized
}
