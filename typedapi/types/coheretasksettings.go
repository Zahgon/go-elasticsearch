package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cohereinputtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheretruncatetype"
)

type CohereTaskSettings struct {
	InputType cohereinputtype.CohereInputType `json:"input_type"`

	ReturnDocuments *bool `json:"return_documents,omitempty"`

	TopN *int `json:"top_n,omitempty"`

	Truncate *coheretruncatetype.CohereTruncateType `json:"truncate,omitempty"`
}

func (s *CohereTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCohereTaskSettings() *CohereTaskSettings { _ = "STUB: not implemented"; return nil }

type CohereTaskSettingsVariant interface {
	CohereTaskSettingsCaster() *CohereTaskSettings
}

func (s *CohereTaskSettings) CohereTaskSettingsCaster() *CohereTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
