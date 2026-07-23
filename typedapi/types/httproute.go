package types

type HttpRoute struct {
	Requests  HttpRouteRequests  `json:"requests"`
	Responses HttpRouteResponses `json:"responses"`
}

func NewHttpRoute() *HttpRoute { _ = "STUB: not implemented"; return nil }
