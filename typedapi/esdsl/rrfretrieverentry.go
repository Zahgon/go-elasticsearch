package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rRFRetrieverEntry struct {
	v types.RRFRetrieverEntry
}

func NewRRFRetrieverEntry() *_rRFRetrieverEntry { _ = "STUB: not implemented"; return nil }

func (u *_rRFRetrieverEntry) RetrieverContainer(retrievercontainer types.RetrieverContainerVariant) *_rRFRetrieverEntry {
	_ = "STUB: not implemented"
	return nil
}

func (u *_retrieverContainer) RRFRetrieverEntryCaster() *types.RRFRetrieverEntry {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rRFRetrieverEntry) RRFRetrieverComponent(rrfretrievercomponent types.RRFRetrieverComponentVariant) *_rRFRetrieverEntry {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rRFRetrieverComponent) RRFRetrieverEntryCaster() *types.RRFRetrieverEntry {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rRFRetrieverEntry) RRFRetrieverEntryCaster() *types.RRFRetrieverEntry {
	_ = "STUB: not implemented"
	return nil
}
