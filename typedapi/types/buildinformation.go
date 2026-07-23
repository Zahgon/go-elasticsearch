package types

type BuildInformation struct {
	Date DateTime `json:"date"`
	Hash string   `json:"hash"`
}

func (s *BuildInformation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBuildInformation() *BuildInformation { _ = "STUB: not implemented"; return nil }
