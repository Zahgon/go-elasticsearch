package types

type Ensemble struct {
	AggregateOutput      *AggregateOutput `json:"aggregate_output,omitempty"`
	ClassificationLabels []string         `json:"classification_labels,omitempty"`
	FeatureNames         []string         `json:"feature_names,omitempty"`
	TargetType           *string          `json:"target_type,omitempty"`
	TrainedModels        []TrainedModel   `json:"trained_models"`
}

func (s *Ensemble) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEnsemble() *Ensemble { _ = "STUB: not implemented"; return nil }

type EnsembleVariant interface {
	EnsembleCaster() *Ensemble
}

func (s *Ensemble) EnsembleCaster() *Ensemble { _ = "STUB: not implemented"; return nil }
