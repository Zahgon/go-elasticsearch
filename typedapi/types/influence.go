package types

type Influence struct {
	InfluencerFieldName   string   `json:"influencer_field_name"`
	InfluencerFieldValues []string `json:"influencer_field_values"`
}

func (s *Influence) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInfluence() *Influence { _ = "STUB: not implemented"; return nil }
