package types

type EqlFeaturesPipes struct {
	PipeHead uint `json:"pipe_head"`
	PipeTail uint `json:"pipe_tail"`
}

func NewEqlFeaturesPipes() *EqlFeaturesPipes { _ = "STUB: not implemented"; return nil }
