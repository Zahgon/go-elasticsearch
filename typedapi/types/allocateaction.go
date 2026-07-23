package types

type AllocateAction struct {
	Exclude            map[string]string `json:"exclude,omitempty"`
	Include            map[string]string `json:"include,omitempty"`
	NumberOfReplicas   *int              `json:"number_of_replicas,omitempty"`
	Require            map[string]string `json:"require,omitempty"`
	TotalShardsPerNode *int              `json:"total_shards_per_node,omitempty"`
}

func (s *AllocateAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAllocateAction() *AllocateAction { _ = "STUB: not implemented"; return nil }

type AllocateActionVariant interface {
	AllocateActionCaster() *AllocateAction
}

func (s *AllocateAction) AllocateActionCaster() *AllocateAction {
	_ = "STUB: not implemented"
	return nil
}
