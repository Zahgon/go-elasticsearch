package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _likeDocument struct {
	v *types.LikeDocument
}

func NewLikeDocument() *_likeDocument { _ = "STUB: not implemented"; return nil }

func (s *_likeDocument) Doc(doc json.RawMessage) *_likeDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_likeDocument) Fields(fields ...string) *_likeDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_likeDocument) Id_(id string) *_likeDocument { _ = "STUB: not implemented"; return nil }

func (s *_likeDocument) Index_(indexname string) *_likeDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_likeDocument) PerFieldAnalyzer(perfieldanalyzer map[string]string) *_likeDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_likeDocument) AddPerFieldAnalyzer(key string, value string) *_likeDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_likeDocument) Routing(routings ...string) *_likeDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_likeDocument) Version(versionnumber int64) *_likeDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_likeDocument) VersionType(versiontype versiontype.VersionType) *_likeDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_likeDocument) LikeDocumentCaster() *types.LikeDocument {
	_ = "STUB: not implemented"
	return nil
}
