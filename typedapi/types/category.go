package types

type Category struct {
	CategoryId uint64 `json:"category_id"`

	Examples []string `json:"examples"`

	GrokPattern *string `json:"grok_pattern,omitempty"`

	JobId string `json:"job_id"`

	MaxMatchingLength uint64 `json:"max_matching_length"`
	Mlcategory        string `json:"mlcategory"`

	NumMatches *int64  `json:"num_matches,omitempty"`
	P          *string `json:"p,omitempty"`

	PartitionFieldName *string `json:"partition_field_name,omitempty"`

	PartitionFieldValue *string `json:"partition_field_value,omitempty"`

	PreferredToCategories []string `json:"preferred_to_categories,omitempty"`

	Regex      string `json:"regex"`
	ResultType string `json:"result_type"`

	Terms string `json:"terms"`
}

func (s *Category) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCategory() *Category { _ = "STUB: not implemented"; return nil }
