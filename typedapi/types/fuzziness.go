package types

type Fuzziness any

type FuzzinessVariant interface {
	FuzzinessCaster() *Fuzziness
}
