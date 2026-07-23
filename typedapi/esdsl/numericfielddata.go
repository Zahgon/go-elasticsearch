package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/numericfielddataformat"
)

type _numericFielddata struct {
	v *types.NumericFielddata
}

func NewNumericFielddata(format numericfielddataformat.NumericFielddataFormat) *_numericFielddata {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numericFielddata) Format(format numericfielddataformat.NumericFielddataFormat) *_numericFielddata {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numericFielddata) NumericFielddataCaster() *types.NumericFielddata {
	_ = "STUB: not implemented"
	return nil
}
