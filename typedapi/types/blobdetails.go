package types

type BlobDetails struct {
	Name string `json:"name"`

	Overwritten bool `json:"overwritten"`
	ReadEarly   bool `json:"read_early"`

	ReadEnd int64 `json:"read_end"`

	ReadStart int64 `json:"read_start"`

	Reads ReadBlobDetails `json:"reads"`

	Size ByteSize `json:"size"`

	SizeBytes int64 `json:"size_bytes"`
}

func (s *BlobDetails) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBlobDetails() *BlobDetails { _ = "STUB: not implemented"; return nil }
