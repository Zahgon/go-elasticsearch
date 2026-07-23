package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaitextembeddingtask"
)

type JinaAITaskSettings struct {
	InputType *jinaaitextembeddingtask.JinaAITextEmbeddingTask `json:"input_type,omitempty"`

	LateChunking *bool `json:"late_chunking,omitempty"`

	ReturnDocuments *bool `json:"return_documents,omitempty"`

	TopN *int `json:"top_n,omitempty"`
}

func (s *JinaAITaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewJinaAITaskSettings() *JinaAITaskSettings { _ = "STUB: not implemented"; return nil }

type JinaAITaskSettingsVariant interface {
	JinaAITaskSettingsCaster() *JinaAITaskSettings
}

func (s *JinaAITaskSettings) JinaAITaskSettingsCaster() *JinaAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
