package types

type DataStreams struct {
	Available    bool  `json:"available"`
	DataStreams  int64 `json:"data_streams"`
	Enabled      bool  `json:"enabled"`
	IndicesCount int64 `json:"indices_count"`
}

func (s *DataStreams) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDataStreams() *DataStreams { _ = "STUB: not implemented"; return nil }
