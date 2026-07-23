package types

type ShardPath struct {
	DataPath         string `json:"data_path"`
	IsCustomDataPath bool   `json:"is_custom_data_path"`
	StatePath        string `json:"state_path"`
}

func (s *ShardPath) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardPath() *ShardPath { _ = "STUB: not implemented"; return nil }
