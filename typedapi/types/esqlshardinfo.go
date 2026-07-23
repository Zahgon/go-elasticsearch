package types

type EsqlShardInfo struct {
	Failed     *int `json:"failed,omitempty"`
	Skipped    *int `json:"skipped,omitempty"`
	Successful *int `json:"successful,omitempty"`
	Total      int  `json:"total"`
}

func (s *EsqlShardInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEsqlShardInfo() *EsqlShardInfo { _ = "STUB: not implemented"; return nil }
