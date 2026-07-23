package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _reindexDestination struct {
	v *types.ReindexDestination
}

func NewReindexDestination() *_reindexDestination { _ = "STUB: not implemented"; return nil }

func (s *_reindexDestination) Index(indexname string) *_reindexDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexDestination) OpType(optype optype.OpType) *_reindexDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexDestination) Pipeline(pipeline string) *_reindexDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexDestination) Routing(routing string) *_reindexDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexDestination) VersionType(versiontype versiontype.VersionType) *_reindexDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexDestination) ReindexDestinationCaster() *types.ReindexDestination {
	_ = "STUB: not implemented"
	return nil
}
