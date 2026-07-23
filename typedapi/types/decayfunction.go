package types

type DecayFunction any

type DecayFunctionVariant interface {
	DecayFunctionCaster() *DecayFunction
}
