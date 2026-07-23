package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type DocumentSimulation struct {
	DocumentSimulation map[string]string `json:"-"`

	Id_ string `json:"_id"`

	Index_  string `json:"_index"`
	Ingest_ Ingest `json:"_ingest"`

	Routing_ *string `json:"_routing,omitempty"`

	Source_      map[string]json.RawMessage `json:"_source"`
	VersionType_ *versiontype.VersionType   `json:"_version_type,omitempty"`
	Version_     StringifiedVersionNumber   `json:"_version,omitempty"`
}

func (s *DocumentSimulation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s DocumentSimulation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDocumentSimulation() *DocumentSimulation { _ = "STUB: not implemented"; return nil }
