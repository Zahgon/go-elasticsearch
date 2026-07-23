package types

type ExtendedSearchUsage struct {
	Retrievers *ExtendedRetrieversSearchUsage `json:"retrievers,omitempty"`
	Section    *ExtendedSectionSearchUsage    `json:"section,omitempty"`
}

func NewExtendedSearchUsage() *ExtendedSearchUsage { _ = "STUB: not implemented"; return nil }
