package types

type MlInference struct {
	Deployments      *MlInferenceDeployments               `json:"deployments,omitempty"`
	IngestProcessors map[string]MlInferenceIngestProcessor `json:"ingest_processors"`
	TrainedModels    MlInferenceTrainedModels              `json:"trained_models"`
}

func NewMlInference() *MlInference { _ = "STUB: not implemented"; return nil }
