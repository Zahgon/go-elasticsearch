package types

type ForceMergeAction struct {
	IndexCodec     *string `json:"index_codec,omitempty"`
	MaxNumSegments int     `json:"max_num_segments"`
}

func (s *ForceMergeAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewForceMergeAction() *ForceMergeAction { _ = "STUB: not implemented"; return nil }

type ForceMergeActionVariant interface {
	ForceMergeActionCaster() *ForceMergeAction
}

func (s *ForceMergeAction) ForceMergeActionCaster() *ForceMergeAction {
	_ = "STUB: not implemented"
	return nil
}
