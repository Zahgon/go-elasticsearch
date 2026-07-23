package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type Script struct {
	Id *string `json:"id,omitempty"`

	Lang    *scriptlanguage.ScriptLanguage `json:"lang,omitempty"`
	Options map[string]string              `json:"options,omitempty"`

	Params map[string]json.RawMessage `json:"params,omitempty"`

	Source ScriptSource `json:"source,omitempty"`
}

func (s *Script) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScript() *Script { _ = "STUB: not implemented"; return nil }

type ScriptVariant interface {
	ScriptCaster() *Script
}

func (s *Script) ScriptCaster() *Script { _ = "STUB: not implemented"; return nil }
