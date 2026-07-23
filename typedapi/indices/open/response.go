package open

type Response struct {
	Acknowledged       bool `json:"acknowledged"`
	ShardsAcknowledged bool `json:"shards_acknowledged"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
