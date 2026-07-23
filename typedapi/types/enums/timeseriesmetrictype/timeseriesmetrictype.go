package timeseriesmetrictype

type TimeSeriesMetricType struct {
	Name string
}

var (
	Gauge = TimeSeriesMetricType{"gauge"}

	Counter = TimeSeriesMetricType{"counter"}

	Summary = TimeSeriesMetricType{"summary"}

	Histogram = TimeSeriesMetricType{"histogram"}

	Position = TimeSeriesMetricType{"position"}
)

func (t TimeSeriesMetricType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TimeSeriesMetricType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TimeSeriesMetricType) String() string { _ = "STUB: not implemented"; return "" }
