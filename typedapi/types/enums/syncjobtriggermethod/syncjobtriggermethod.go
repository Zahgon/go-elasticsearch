package syncjobtriggermethod

type SyncJobTriggerMethod struct {
	Name string
}

var (
	Ondemand = SyncJobTriggerMethod{"on_demand"}

	Scheduled = SyncJobTriggerMethod{"scheduled"}
)

func (s SyncJobTriggerMethod) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SyncJobTriggerMethod) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SyncJobTriggerMethod) String() string { _ = "STUB: not implemented"; return "" }
