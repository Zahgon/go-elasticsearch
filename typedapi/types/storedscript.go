package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type StoredScript struct {
	Lang    scriptlanguage.ScriptLanguage `json:"lang"`
	Options map[string]string             `json:"options,omitempty"`

	Source ScriptSource `json:"source"`
}

func (s *StoredScript) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStoredScript() *StoredScript { _ = "STUB: not implemented"; return nil }

type StoredScriptVariant interface {
	StoredScriptCaster() *StoredScript
}

func (s *StoredScript) StoredScriptCaster() *StoredScript { _ = "STUB: not implemented"; return nil }
