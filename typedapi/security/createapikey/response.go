package createapikey

type Response struct {
	ApiKey string `json:"api_key"`

	Encoded string `json:"encoded"`

	Expiration *int64 `json:"expiration,omitempty"`

	Id string `json:"id"`

	Name string `json:"name"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
