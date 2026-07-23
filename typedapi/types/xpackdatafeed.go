package types

type XpackDatafeed struct {
	Count int64 `json:"count"`
}

func (s *XpackDatafeed) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewXpackDatafeed() *XpackDatafeed { _ = "STUB: not implemented"; return nil }
