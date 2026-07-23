package rollover

type Response struct {
	Acknowledged       bool            `json:"acknowledged"`
	Conditions         map[string]bool `json:"conditions"`
	DryRun             bool            `json:"dry_run"`
	NewIndex           string          `json:"new_index"`
	OldIndex           string          `json:"old_index"`
	RolledOver         bool            `json:"rolled_over"`
	ShardsAcknowledged bool            `json:"shards_acknowledged"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
