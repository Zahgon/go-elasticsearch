package types

type HttpRouteResponses struct {
	Count                 int64               `json:"count"`
	HandlingTimeHistogram []TimeHttpHistogram `json:"handling_time_histogram"`
	SizeHistogram         []SizeHttpHistogram `json:"size_histogram"`
	TotalSizeInBytes      int64               `json:"total_size_in_bytes"`
}

func (s *HttpRouteResponses) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHttpRouteResponses() *HttpRouteResponses { _ = "STUB: not implemented"; return nil }
