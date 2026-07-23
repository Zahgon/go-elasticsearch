package cardinalityexecutionmode

type CardinalityExecutionMode struct {
	Name string
}

var (
	Globalordinals = CardinalityExecutionMode{"global_ordinals"}

	Segmentordinals = CardinalityExecutionMode{"segment_ordinals"}

	Direct = CardinalityExecutionMode{"direct"}

	Savememoryheuristic = CardinalityExecutionMode{"save_memory_heuristic"}

	Savetimeheuristic = CardinalityExecutionMode{"save_time_heuristic"}
)

func (c CardinalityExecutionMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CardinalityExecutionMode) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CardinalityExecutionMode) String() string { _ = "STUB: not implemented"; return "" }
