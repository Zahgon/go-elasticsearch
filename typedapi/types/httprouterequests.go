package types

type HttpRouteRequests struct {
	Count            int64               `json:"count"`
	SizeHistogram    []SizeHttpHistogram `json:"size_histogram"`
	TotalSizeInBytes int64               `json:"total_size_in_bytes"`
}

func (s *HttpRouteRequests) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHttpRouteRequests() *HttpRouteRequests { _ = "STUB: not implemented"; return nil }
