package types

type RecoveryStatus struct {
	Shards []ShardRecovery `json:"shards"`
}

func NewRecoveryStatus() *RecoveryStatus { _ = "STUB: not implemented"; return nil }
