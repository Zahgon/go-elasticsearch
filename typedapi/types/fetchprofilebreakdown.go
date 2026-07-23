package types

type FetchProfileBreakdown struct {
	LoadSource            *int `json:"load_source,omitempty"`
	LoadSourceCount       *int `json:"load_source_count,omitempty"`
	LoadStoredFields      *int `json:"load_stored_fields,omitempty"`
	LoadStoredFieldsCount *int `json:"load_stored_fields_count,omitempty"`
	NextReader            *int `json:"next_reader,omitempty"`
	NextReaderCount       *int `json:"next_reader_count,omitempty"`
	Process               *int `json:"process,omitempty"`
	ProcessCount          *int `json:"process_count,omitempty"`
}

func (s *FetchProfileBreakdown) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFetchProfileBreakdown() *FetchProfileBreakdown { _ = "STUB: not implemented"; return nil }
