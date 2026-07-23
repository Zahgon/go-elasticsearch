package decision

type Decision struct {
	Name string
}

var (
	Yes = Decision{"yes"}

	No = Decision{"no"}

	Worsebalance = Decision{"worse_balance"}

	Throttled = Decision{"throttled"}

	Awaitinginfo = Decision{"awaiting_info"}

	Allocationdelayed = Decision{"allocation_delayed"}

	Novalidshardcopy = Decision{"no_valid_shard_copy"}

	Noattempt = Decision{"no_attempt"}
)

func (d Decision) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Decision) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d Decision) String() string { _ = "STUB: not implemented"; return "" }
