package types

type MlInferenceIngestProcessor struct {
	NumDocsProcessed MlInferenceIngestProcessorCount `json:"num_docs_processed"`
	NumFailures      MlInferenceIngestProcessorCount `json:"num_failures"`
	Pipelines        MlCounter                       `json:"pipelines"`
	TimeMs           MlInferenceIngestProcessorCount `json:"time_ms"`
}

func NewMlInferenceIngestProcessor() *MlInferenceIngestProcessor {
	_ = "STUB: not implemented"
	return nil
}
