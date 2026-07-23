package types

type CompositeAggregateKey map[string]FieldValue

type CompositeAggregateKeyVariant interface {
	CompositeAggregateKeyCaster() *CompositeAggregateKey
}
