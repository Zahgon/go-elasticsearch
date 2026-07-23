package upgradetransforms

type Response struct {
	NeedsUpdate int `json:"needs_update"`

	NoAction int `json:"no_action"`

	Updated int `json:"updated"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
