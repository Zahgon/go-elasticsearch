package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _painlessContextSetup struct {
	v *types.PainlessContextSetup
}

func NewPainlessContextSetup(document json.RawMessage) *_painlessContextSetup {
	_ = "STUB: not implemented"
	return nil
}

func (s *_painlessContextSetup) Document(document json.RawMessage) *_painlessContextSetup {
	_ = "STUB: not implemented"
	return nil
}

func (s *_painlessContextSetup) Index(indexname string) *_painlessContextSetup {
	_ = "STUB: not implemented"
	return nil
}

func (s *_painlessContextSetup) Query(query types.QueryVariant) *_painlessContextSetup {
	_ = "STUB: not implemented"
	return nil
}

func (s *_painlessContextSetup) PainlessContextSetupCaster() *types.PainlessContextSetup {
	_ = "STUB: not implemented"
	return nil
}
