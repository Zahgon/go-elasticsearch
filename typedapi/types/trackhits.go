package types

type TrackHits any

type TrackHitsVariant interface {
	TrackHitsCaster() *TrackHits
}
