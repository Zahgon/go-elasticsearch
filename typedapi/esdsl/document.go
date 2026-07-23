package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _document struct {
	v *types.Document
}

func NewDocument(source_ json.RawMessage) *_document { _ = "STUB: not implemented"; return nil }

func (s *_document) Id_(id string) *_document { _ = "STUB: not implemented"; return nil }

func (s *_document) Index_(indexname string) *_document { _ = "STUB: not implemented"; return nil }

func (s *_document) Source_(source_ json.RawMessage) *_document {
	_ = "STUB: not implemented"
	return nil
}

func (s *_document) DocumentCaster() *types.Document { _ = "STUB: not implemented"; return nil }
