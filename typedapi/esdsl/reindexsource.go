package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _reindexSource struct {
	v *types.ReindexSource
}

func NewReindexSource() *_reindexSource { _ = "STUB: not implemented"; return nil }

func (s *_reindexSource) Index(indices ...string) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) ProjectRouting(projectrouting string) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) Query(query types.QueryVariant) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) Remote(remote types.RemoteSourceVariant) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) Size(size int) *_reindexSource { _ = "STUB: not implemented"; return nil }

func (s *_reindexSource) Slice(slice types.SlicedScrollVariant) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) Sort(sorts ...types.SortCombinationsVariant) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) SortValues(sortvalues []types.SortCombinations) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) SourceFields_(sourceconfig types.SourceConfigVariant) *_reindexSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reindexSource) ReindexSourceCaster() *types.ReindexSource {
	_ = "STUB: not implemented"
	return nil
}
