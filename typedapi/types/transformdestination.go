package types

type TransformDestination struct {
	Index *string `json:"index,omitempty"`

	Pipeline *string `json:"pipeline,omitempty"`
}

func (s *TransformDestination) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTransformDestination() *TransformDestination { _ = "STUB: not implemented"; return nil }

type TransformDestinationVariant interface {
	TransformDestinationCaster() *TransformDestination
}

func (s *TransformDestination) TransformDestinationCaster() *TransformDestination {
	_ = "STUB: not implemented"
	return nil
}
