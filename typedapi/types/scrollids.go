package types

type ScrollIds []string

type ScrollIdsVariant interface {
	ScrollIdsCaster() *ScrollIds
}
