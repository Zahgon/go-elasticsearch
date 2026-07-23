package types

type Percentage any

type PercentageVariant interface {
	PercentageCaster() *Percentage
}
