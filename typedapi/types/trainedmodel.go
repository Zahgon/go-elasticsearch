package types

type TrainedModel struct {
	Ensemble *Ensemble `json:"ensemble,omitempty"`

	Tree *TrainedModelTree `json:"tree,omitempty"`

	TreeNode *TrainedModelTreeNode `json:"tree_node,omitempty"`
}

func NewTrainedModel() *TrainedModel { _ = "STUB: not implemented"; return nil }

type TrainedModelVariant interface {
	TrainedModelCaster() *TrainedModel
}

func (s *TrainedModel) TrainedModelCaster() *TrainedModel { _ = "STUB: not implemented"; return nil }
