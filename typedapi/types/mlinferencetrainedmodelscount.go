package types

type MlInferenceTrainedModelsCount struct {
	Classification *int64 `json:"classification,omitempty"`
	Ner            *int64 `json:"ner,omitempty"`
	Other          int64  `json:"other"`
	PassThrough    *int64 `json:"pass_through,omitempty"`
	Prepackaged    int64  `json:"prepackaged"`
	Regression     *int64 `json:"regression,omitempty"`
	TextEmbedding  *int64 `json:"text_embedding,omitempty"`
	Total          int64  `json:"total"`
}

func (s *MlInferenceTrainedModelsCount) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMlInferenceTrainedModelsCount() *MlInferenceTrainedModelsCount {
	_ = "STUB: not implemented"
	return nil
}
