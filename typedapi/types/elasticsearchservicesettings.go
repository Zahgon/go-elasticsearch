package types

type ElasticsearchServiceSettings struct {
	AdaptiveAllocations *AdaptiveAllocations `json:"adaptive_allocations,omitempty"`

	DeploymentId *string `json:"deployment_id,omitempty"`

	LongDocumentStrategy *string `json:"long_document_strategy,omitempty"`

	MaxChunksPerDoc *int `json:"max_chunks_per_doc,omitempty"`

	ModelId string `json:"model_id"`

	NumAllocations *int `json:"num_allocations,omitempty"`

	NumThreads int `json:"num_threads"`
}

func (s *ElasticsearchServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewElasticsearchServiceSettings() *ElasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type ElasticsearchServiceSettingsVariant interface {
	ElasticsearchServiceSettingsCaster() *ElasticsearchServiceSettings
}

func (s *ElasticsearchServiceSettings) ElasticsearchServiceSettingsCaster() *ElasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
