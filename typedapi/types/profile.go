package types

type Profile struct {
	Request *SearchRequestCoordinatorMetadata `json:"request,omitempty"`
	Shards  []ShardProfile                    `json:"shards"`
}

func NewProfile() *Profile { _ = "STUB: not implemented"; return nil }
