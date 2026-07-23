package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaitextembeddingtask"
)

type _jinaAITaskSettings struct {
	v *types.JinaAITaskSettings
}

func NewJinaAITaskSettings() *_jinaAITaskSettings { _ = "STUB: not implemented"; return nil }

func (s *_jinaAITaskSettings) InputType(inputtype jinaaitextembeddingtask.JinaAITextEmbeddingTask) *_jinaAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAITaskSettings) LateChunking(latechunking bool) *_jinaAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAITaskSettings) ReturnDocuments(returndocuments bool) *_jinaAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAITaskSettings) TopN(topn int) *_jinaAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAITaskSettings) JinaAITaskSettingsCaster() *types.JinaAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
