package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _detectorUpdate struct {
	v *types.DetectorUpdate
}

func NewDetectorUpdate(detectorindex int) *_detectorUpdate { _ = "STUB: not implemented"; return nil }

func (s *_detectorUpdate) CustomRules(customrules ...types.DetectionRuleVariant) *_detectorUpdate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectorUpdate) CustomRulesValues(customrulesvalues []types.DetectionRule) *_detectorUpdate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectorUpdate) Description(description string) *_detectorUpdate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectorUpdate) DetectorIndex(detectorindex int) *_detectorUpdate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectorUpdate) DetectorUpdateCaster() *types.DetectorUpdate {
	_ = "STUB: not implemented"
	return nil
}
