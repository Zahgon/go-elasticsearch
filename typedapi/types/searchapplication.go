package types

type SearchApplication struct {
	AnalyticsCollectionName *string `json:"analytics_collection_name,omitempty"`

	Indices []string `json:"indices"`

	Name string `json:"name"`

	Template *SearchApplicationTemplate `json:"template,omitempty"`

	UpdatedAtMillis int64 `json:"updated_at_millis"`
}

func (s *SearchApplication) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchApplication() *SearchApplication { _ = "STUB: not implemented"; return nil }
