package types

type IngestPipelineParams struct {
	ExtractBinaryContent bool   `json:"extract_binary_content"`
	Name                 string `json:"name"`
	ReduceWhitespace     bool   `json:"reduce_whitespace"`
	RunMlInference       bool   `json:"run_ml_inference"`
}

func (s *IngestPipelineParams) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIngestPipelineParams() *IngestPipelineParams { _ = "STUB: not implemented"; return nil }

type IngestPipelineParamsVariant interface {
	IngestPipelineParamsCaster() *IngestPipelineParams
}

func (s *IngestPipelineParams) IngestPipelineParamsCaster() *IngestPipelineParams {
	_ = "STUB: not implemented"
	return nil
}
