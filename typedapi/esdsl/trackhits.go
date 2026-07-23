package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trackHits struct {
	v types.TrackHits
}

func NewTrackHits() *_trackHits { _ = "STUB: not implemented"; return nil }

func (u *_trackHits) Bool(bool bool) *_trackHits { _ = "STUB: not implemented"; return nil }

func (u *_trackHits) Int(int int) *_trackHits { _ = "STUB: not implemented"; return nil }

func (u *_trackHits) TrackHitsCaster() *types.TrackHits { _ = "STUB: not implemented"; return nil }
