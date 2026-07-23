package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _epochTimeUnitSeconds struct {
	v types.EpochTimeUnitSeconds
}

func NewEpochTimeUnitSeconds(epochtimeunitseconds int64) *_epochTimeUnitSeconds {
	_ = "STUB: not implemented"
	return nil
}

func (u *_epochTimeUnitSeconds) EpochTimeUnitSecondsCaster() *types.EpochTimeUnitSeconds {
	_ = "STUB: not implemented"
	return nil
}
