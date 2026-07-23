package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/numericfielddataformat"
)

type NumericFielddata struct {
	Format numericfielddataformat.NumericFielddataFormat `json:"format"`
}

func NewNumericFielddata() *NumericFielddata { _ = "STUB: not implemented"; return nil }

type NumericFielddataVariant interface {
	NumericFielddataCaster() *NumericFielddata
}

func (s *NumericFielddata) NumericFielddataCaster() *NumericFielddata {
	_ = "STUB: not implemented"
	return nil
}
