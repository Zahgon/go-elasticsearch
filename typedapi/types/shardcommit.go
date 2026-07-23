package types

type ShardCommit struct {
	Generation int               `json:"generation"`
	Id         string            `json:"id"`
	NumDocs    int64             `json:"num_docs"`
	UserData   map[string]string `json:"user_data"`
}

func (s *ShardCommit) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardCommit() *ShardCommit { _ = "STUB: not implemented"; return nil }
