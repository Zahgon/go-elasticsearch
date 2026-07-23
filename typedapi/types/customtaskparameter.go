package types

type CustomTaskParameter any

type CustomTaskParameterVariant interface {
	CustomTaskParameterCaster() *CustomTaskParameter
}
