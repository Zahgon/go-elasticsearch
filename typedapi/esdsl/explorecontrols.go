package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _exploreControls struct {
	v *types.ExploreControls
}

func NewExploreControls(usesignificance bool) *_exploreControls {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exploreControls) SampleDiversity(samplediversity types.SampleDiversityVariant) *_exploreControls {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exploreControls) SampleSize(samplesize int) *_exploreControls {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exploreControls) Timeout(duration types.DurationVariant) *_exploreControls {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exploreControls) UseSignificance(usesignificance bool) *_exploreControls {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exploreControls) ExploreControlsCaster() *types.ExploreControls {
	_ = "STUB: not implemented"
	return nil
}
