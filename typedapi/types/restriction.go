package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/restrictionworkflow"
)

type Restriction struct {
	Workflows []restrictionworkflow.RestrictionWorkflow `json:"workflows"`
}

func NewRestriction() *Restriction { _ = "STUB: not implemented"; return nil }

type RestrictionVariant interface {
	RestrictionCaster() *Restriction
}

func (s *Restriction) RestrictionCaster() *Restriction { _ = "STUB: not implemented"; return nil }
