package putwatch

type Response struct {
	Created      bool   `json:"created"`
	Id_          string `json:"_id"`
	PrimaryTerm_ int64  `json:"_primary_term"`
	SeqNo_       int64  `json:"_seq_no"`
	Version_     int64  `json:"_version"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
