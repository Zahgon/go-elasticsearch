package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fileContent struct {
	v *types.FileContent
}

func NewFileContent(filedata string, filename string) *_fileContent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fileContent) FileData(filedata string) *_fileContent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fileContent) Filename(filename string) *_fileContent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fileContent) FileContentCaster() *types.FileContent {
	_ = "STUB: not implemented"
	return nil
}
