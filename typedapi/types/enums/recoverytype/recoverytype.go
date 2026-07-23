package recoverytype

type RecoveryType struct {
	Name string
}

var (
	EMPTYSTORE = RecoveryType{"EMPTY_STORE"}

	EXISTINGSTORE = RecoveryType{"EXISTING_STORE"}

	LOCALSHARDS = RecoveryType{"LOCAL_SHARDS"}

	PEER = RecoveryType{"PEER"}

	SNAPSHOT = RecoveryType{"SNAPSHOT"}
)

func (r RecoveryType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RecoveryType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RecoveryType) String() string { _ = "STUB: not implemented"; return "" }
