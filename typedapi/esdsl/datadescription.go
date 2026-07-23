package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataDescription struct {
	v *types.DataDescription
}

func NewDataDescription() *_dataDescription { _ = "STUB: not implemented"; return nil }

func (s *_dataDescription) FieldDelimiter(fielddelimiter string) *_dataDescription {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataDescription) Format(format string) *_dataDescription {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataDescription) TimeField(field string) *_dataDescription {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataDescription) TimeFormat(timeformat string) *_dataDescription {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataDescription) DataDescriptionCaster() *types.DataDescription {
	_ = "STUB: not implemented"
	return nil
}
