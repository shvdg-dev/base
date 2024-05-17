package app

import "base/pkg/i18n"

var I18n *i18n.I18n

func init() {
	I18n = i18n.NewConfig()
}
