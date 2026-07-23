package types

type GeoHashPrecision any

type GeoHashPrecisionVariant interface {
	GeoHashPrecisionCaster() *GeoHashPrecision
}
