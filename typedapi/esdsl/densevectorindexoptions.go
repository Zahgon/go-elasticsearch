package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorindexoptionstype"
)

type _denseVectorIndexOptions struct {
	v *types.DenseVectorIndexOptions
}

func NewDenseVectorIndexOptions(type_ densevectorindexoptionstype.DenseVectorIndexOptionsType) *_denseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorIndexOptions) ConfidenceInterval(confidenceinterval float32) *_denseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorIndexOptions) EfConstruction(efconstruction int) *_denseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorIndexOptions) FlatIndexThreshold(flatindexthreshold int) *_denseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorIndexOptions) M(m int) *_denseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorIndexOptions) OnDiskRescore(ondiskrescore bool) *_denseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorIndexOptions) RescoreVector(rescorevector types.DenseVectorIndexOptionsRescoreVectorVariant) *_denseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorIndexOptions) Type(type_ densevectorindexoptionstype.DenseVectorIndexOptionsType) *_denseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorIndexOptions) DenseVectorIndexOptionsCaster() *types.DenseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}
