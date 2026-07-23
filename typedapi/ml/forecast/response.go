package forecast

type Response struct {
	Acknowledged bool   `json:"acknowledged"`
	ForecastId   string `json:"forecast_id"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
