package types

type TrainedModelTree struct {
	ClassificationLabels []string               `json:"classification_labels,omitempty"`
	FeatureNames         []string               `json:"feature_names"`
	TargetType           *string                `json:"target_type,omitempty"`
	TreeStructure        []TrainedModelTreeNode `json:"tree_structure"`
}

func (s *TrainedModelTree) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTrainedModelTree() *TrainedModelTree { _ = "STUB: not implemented"; return nil }

type TrainedModelTreeVariant interface {
	TrainedModelTreeCaster() *TrainedModelTree
}

func (s *TrainedModelTree) TrainedModelTreeCaster() *TrainedModelTree {
	_ = "STUB: not implemented"
	return nil
}
