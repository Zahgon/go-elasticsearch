package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type ScriptProcessor struct {
	Description *string `json:"description,omitempty"`

	Id *string `json:"id,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	Lang *scriptlanguage.ScriptLanguage `json:"lang,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Params map[string]json.RawMessage `json:"params,omitempty"`

	Source ScriptSource `json:"source,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *ScriptProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScriptProcessor() *ScriptProcessor { _ = "STUB: not implemented"; return nil }

type ScriptProcessorVariant interface {
	ScriptProcessorCaster() *ScriptProcessor
}

func (s *ScriptProcessor) ScriptProcessorCaster() *ScriptProcessor {
	_ = "STUB: not implemented"
	return nil
}
