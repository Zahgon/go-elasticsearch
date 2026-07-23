package types

type InferenceStringGroup []InferenceString

type InferenceStringGroupVariant interface {
	InferenceStringGroupCaster() *InferenceStringGroup
}
