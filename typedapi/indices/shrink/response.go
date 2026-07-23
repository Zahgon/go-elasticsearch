package shrink

type Response struct {
	Acknowledged       bool   `json:"acknowledged"`
	Index              string `json:"index"`
	ShardsAcknowledged bool   `json:"shards_acknowledged"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
