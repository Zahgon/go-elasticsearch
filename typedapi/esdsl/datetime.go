package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dateTime struct {
	v types.DateTime
}

func NewDateTime() *_dateTime { _ = "STUB: not implemented"; return nil }

func (u *_dateTime) String(string string) *_dateTime { _ = "STUB: not implemented"; return nil }

func (u *_dateTime) EpochTimeUnitMillis(epochtimeunitmillis int64) *_dateTime {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dateTime) DateTimeCaster() *types.DateTime { _ = "STUB: not implemented"; return nil }
