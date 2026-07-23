package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _requestEmbedding struct {
	v *types.RequestEmbedding
}

func NewRequestEmbedding() *_requestEmbedding { _ = "STUB: not implemented"; return nil }

func (s *_requestEmbedding) Input(embeddinginput types.EmbeddingInputVariant) *_requestEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestEmbedding) InputType(inputtype string) *_requestEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestEmbedding) TaskSettings(tasksettings json.RawMessage) *_requestEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestEmbedding) RequestEmbeddingCaster() *types.RequestEmbedding {
	_ = "STUB: not implemented"
	return nil
}
