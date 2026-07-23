package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trainedModelTreeNode struct {
	v *types.TrainedModelTreeNode
}

func NewTrainedModelTreeNode(nodeindex int) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) DecisionType(decisiontype string) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) DefaultLeft(defaultleft bool) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) LeafValue(leafvalue types.Float64) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) LeftChild(leftchild int) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) NodeIndex(nodeindex int) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) RightChild(rightchild int) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) SplitFeature(splitfeature int) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) SplitGain(splitgain int) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) Threshold(threshold types.Float64) *_trainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTreeNode) TrainedModelTreeNodeCaster() *types.TrainedModelTreeNode {
	_ = "STUB: not implemented"
	return nil
}
