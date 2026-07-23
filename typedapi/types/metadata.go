package types

import (
	"encoding/json"
)

type Metadata map[string]json.RawMessage

type MetadataVariant interface {
	MetadataCaster() *Metadata
}
