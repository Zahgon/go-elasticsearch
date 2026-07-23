package types

type TransformHealthIssue struct {
	Count int `json:"count"`

	Details              *string  `json:"details,omitempty"`
	FirstOccurenceString DateTime `json:"first_occurence_string,omitempty"`

	FirstOccurrence *int64 `json:"first_occurrence,omitempty"`

	Issue string `json:"issue"`

	Type string `json:"type"`
}

func (s *TransformHealthIssue) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTransformHealthIssue() *TransformHealthIssue { _ = "STUB: not implemented"; return nil }
