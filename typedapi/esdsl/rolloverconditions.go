package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rolloverConditions struct {
	v *types.RolloverConditions
}

func NewRolloverConditions() *_rolloverConditions { _ = "STUB: not implemented"; return nil }

func (s *_rolloverConditions) MaxAge(duration types.DurationVariant) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MaxAgeMillis(durationvalueunitmillis int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MaxDocs(maxdocs int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MaxPrimaryShardDocs(maxprimarysharddocs int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MaxPrimaryShardSize(bytesize types.ByteSizeVariant) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MaxPrimaryShardSizeBytes(maxprimaryshardsizebytes int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MaxSize(bytesize types.ByteSizeVariant) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MaxSizeBytes(maxsizebytes int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MinAge(duration types.DurationVariant) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MinDocs(mindocs int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MinPrimaryShardDocs(minprimarysharddocs int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MinPrimaryShardSize(bytesize types.ByteSizeVariant) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MinPrimaryShardSizeBytes(minprimaryshardsizebytes int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MinSize(bytesize types.ByteSizeVariant) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) MinSizeBytes(minsizebytes int64) *_rolloverConditions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverConditions) RolloverConditionsCaster() *types.RolloverConditions {
	_ = "STUB: not implemented"
	return nil
}
