package refresh

type Refresh struct {
	Name string
}

var (
	True = Refresh{"true"}

	False = Refresh{"false"}

	Waitfor = Refresh{"wait_for"}
)

func (r *Refresh) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (r Refresh) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Refresh) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r Refresh) String() string { _ = "STUB: not implemented"; return "" }
