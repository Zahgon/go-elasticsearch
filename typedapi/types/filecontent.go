package types

type FileContent struct {
	FileData string `json:"file_data"`

	Filename string `json:"filename"`
}

func (s *FileContent) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFileContent() *FileContent { _ = "STUB: not implemented"; return nil }

type FileContentVariant interface {
	FileContentCaster() *FileContent
}

func (s *FileContent) FileContentCaster() *FileContent { _ = "STUB: not implemented"; return nil }
