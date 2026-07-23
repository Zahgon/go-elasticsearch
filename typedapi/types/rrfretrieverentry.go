package types

type RRFRetrieverEntry any

type RRFRetrieverEntryVariant interface {
	RRFRetrieverEntryCaster() *RRFRetrieverEntry
}
