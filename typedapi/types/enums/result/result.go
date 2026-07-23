package result

type Result struct {
	Name string
}

var (
	Created = Result{"created"}

	Updated = Result{"updated"}

	Deleted = Result{"deleted"}

	Notfound = Result{"not_found"}

	Noop = Result{"noop"}
)

func (r Result) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Result) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r Result) String() string { _ = "STUB: not implemented"; return "" }
