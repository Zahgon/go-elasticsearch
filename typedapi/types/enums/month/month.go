package month

type Month struct {
	Name string
}

var (
	January = Month{"january"}

	February = Month{"february"}

	March = Month{"march"}

	April = Month{"april"}

	May = Month{"may"}

	June = Month{"june"}

	July = Month{"july"}

	August = Month{"august"}

	September = Month{"september"}

	October = Month{"october"}

	November = Month{"november"}

	December = Month{"december"}
)

func (m Month) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Month) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m Month) String() string { _ = "STUB: not implemented"; return "" }
