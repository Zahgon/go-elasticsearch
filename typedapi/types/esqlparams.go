package types

type ESQLParams any

type ESQLParamsVariant interface {
	ESQLParamsCaster() *ESQLParams
}
