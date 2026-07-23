package types

type FileDetails struct {
	Length    int64  `json:"length"`
	Name      string `json:"name"`
	Recovered int64  `json:"recovered"`
}

func (s *FileDetails) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFileDetails() *FileDetails { _ = "STUB: not implemented"; return nil }
