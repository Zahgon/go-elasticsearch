package trainingpriority

type TrainingPriority struct {
	Name string
}

var (
	Normal = TrainingPriority{"normal"}

	Low = TrainingPriority{"low"}
)

func (t TrainingPriority) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TrainingPriority) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TrainingPriority) String() string { _ = "STUB: not implemented"; return "" }
