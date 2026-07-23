package types

type SourceConfig any

type SourceConfigVariant interface {
	SourceConfigCaster() *SourceConfig
}
