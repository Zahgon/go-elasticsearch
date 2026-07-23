package deletewatch

type Response struct {
	Found    bool   `json:"found"`
	Id_      string `json:"_id"`
	Version_ int64  `json:"_version"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
