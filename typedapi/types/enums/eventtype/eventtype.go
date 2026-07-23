package eventtype

type EventType struct {
	Name string
}

var (
	PageView = EventType{"page_view"}

	Search = EventType{"search"}

	SearchClick = EventType{"search_click"}
)

func (e EventType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EventType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e EventType) String() string { _ = "STUB: not implemented"; return "" }
