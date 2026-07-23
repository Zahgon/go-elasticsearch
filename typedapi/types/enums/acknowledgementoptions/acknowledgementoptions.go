package acknowledgementoptions

type AcknowledgementOptions struct {
	Name string
}

var (
	Awaitssuccessfulexecution = AcknowledgementOptions{"awaits_successful_execution"}

	Ackable = AcknowledgementOptions{"ackable"}

	Acked = AcknowledgementOptions{"acked"}
)

func (a AcknowledgementOptions) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AcknowledgementOptions) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AcknowledgementOptions) String() string { _ = "STUB: not implemented"; return "" }
