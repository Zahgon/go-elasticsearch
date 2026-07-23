package getbasicstatus

type Response struct {
	EligibleToStartBasic bool `json:"eligible_to_start_basic"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
