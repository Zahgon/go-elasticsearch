package types

type SourceOnlyRepositorySettings any

type SourceOnlyRepositorySettingsVariant interface {
	SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings
}
