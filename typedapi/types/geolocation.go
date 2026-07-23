package types

type GeoLocation any

type GeoLocationVariant interface {
	GeoLocationCaster() *GeoLocation
}
