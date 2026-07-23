package types

type SimpleQueryStringFlags PipeSeparatedFlagsSimpleQueryStringFlag

type SimpleQueryStringFlagsVariant interface {
	SimpleQueryStringFlagsCaster() *SimpleQueryStringFlags
}
