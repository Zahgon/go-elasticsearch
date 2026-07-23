package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

type ExpandWildcards []expandwildcard.ExpandWildcard

type ExpandWildcardsVariant interface {
	ExpandWildcardsCaster() *ExpandWildcards
}
