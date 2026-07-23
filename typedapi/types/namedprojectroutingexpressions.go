package types

type NamedProjectRoutingExpressions map[string]ProjectRoutingExpression

type NamedProjectRoutingExpressionsVariant interface {
	NamedProjectRoutingExpressionsCaster() *NamedProjectRoutingExpressions
}
