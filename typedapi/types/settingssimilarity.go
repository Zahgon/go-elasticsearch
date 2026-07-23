package types

type SettingsSimilarity any

type SettingsSimilarityVariant interface {
	SettingsSimilarityCaster() *SettingsSimilarity
}
