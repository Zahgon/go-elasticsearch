package recoverystage

type RecoveryStage struct {
	Name string
}

var (
	INIT = RecoveryStage{"INIT"}

	INDEX = RecoveryStage{"INDEX"}

	VERIFYINDEX = RecoveryStage{"VERIFY_INDEX"}

	TRANSLOG = RecoveryStage{"TRANSLOG"}

	FINALIZE = RecoveryStage{"FINALIZE"}

	DONE = RecoveryStage{"DONE"}
)

func (r RecoveryStage) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RecoveryStage) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RecoveryStage) String() string { _ = "STUB: not implemented"; return "" }
