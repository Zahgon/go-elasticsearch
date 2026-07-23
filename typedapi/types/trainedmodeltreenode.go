package types

type TrainedModelTreeNode struct {
	DecisionType *string  `json:"decision_type,omitempty"`
	DefaultLeft  *bool    `json:"default_left,omitempty"`
	LeafValue    *Float64 `json:"leaf_value,omitempty"`
	LeftChild    *int     `json:"left_child,omitempty"`
	NodeIndex    int      `json:"node_index"`
	RightChild   *int     `json:"right_child,omitempty"`
	SplitFeature *int     `json:"split_feature,omitempty"`
	SplitGain    *int     `json:"split_gain,omitempty"`
	Threshold    *Float64 `json:"threshold,omitempty"`
}

func (s *TrainedModelTreeNode) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelTreeNode() *TrainedModelTreeNode { _ = "STUB: not implemented"; return nil }

type TrainedModelTreeNodeVariant interface {
	TrainedModelTreeNodeCaster() *TrainedModelTreeNode
}

func (s *TrainedModelTreeNode) TrainedModelTreeNodeCaster() *TrainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}
