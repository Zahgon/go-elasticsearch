package routingstate

type RoutingState struct {
	Name string
}

var (
	Failed = RoutingState{"failed"}

	Started = RoutingState{"started"}

	Starting = RoutingState{"starting"}

	Stopped = RoutingState{"stopped"}

	Stopping = RoutingState{"stopping"}
)

func (r RoutingState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RoutingState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RoutingState) String() string { _ = "STUB: not implemented"; return "" }
