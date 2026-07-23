package putfilter

type Response struct {
	Description string   `json:"description"`
	FilterId    string   `json:"filter_id"`
	Items       []string `json:"items"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
