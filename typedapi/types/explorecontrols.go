package types

type ExploreControls struct {
	SampleDiversity *SampleDiversity `json:"sample_diversity,omitempty"`

	SampleSize *int `json:"sample_size,omitempty"`

	Timeout Duration `json:"timeout,omitempty"`

	UseSignificance bool `json:"use_significance"`
}

func (s *ExploreControls) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewExploreControls() *ExploreControls { _ = "STUB: not implemented"; return nil }

type ExploreControlsVariant interface {
	ExploreControlsCaster() *ExploreControls
}

func (s *ExploreControls) ExploreControlsCaster() *ExploreControls {
	_ = "STUB: not implemented"
	return nil
}
