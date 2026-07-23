package followerindexstatus

type FollowerIndexStatus struct {
	Name string
}

var (
	Active = FollowerIndexStatus{"active"}

	Paused = FollowerIndexStatus{"paused"}
)

func (f FollowerIndexStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FollowerIndexStatus) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FollowerIndexStatus) String() string { _ = "STUB: not implemented"; return "" }
