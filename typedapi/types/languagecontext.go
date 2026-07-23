package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type LanguageContext struct {
	Contexts []string                      `json:"contexts"`
	Language scriptlanguage.ScriptLanguage `json:"language"`
}

func NewLanguageContext() *LanguageContext { _ = "STUB: not implemented"; return nil }
