package catcircuitbreakercolumn

type CatCircuitBreakerColumn struct {
	Name string
}

var (
	Nodeid = CatCircuitBreakerColumn{"node_id"}

	Nodename = CatCircuitBreakerColumn{"node_name"}

	Breaker = CatCircuitBreakerColumn{"breaker"}

	Limit = CatCircuitBreakerColumn{"limit"}

	Limitbytes = CatCircuitBreakerColumn{"limit_bytes"}

	Estimated = CatCircuitBreakerColumn{"estimated"}

	Estimatedbytes = CatCircuitBreakerColumn{"estimated_bytes"}

	Tripped = CatCircuitBreakerColumn{"tripped"}

	Overhead = CatCircuitBreakerColumn{"overhead"}
)

func (c CatCircuitBreakerColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatCircuitBreakerColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatCircuitBreakerColumn) String() string { _ = "STUB: not implemented"; return "" }
