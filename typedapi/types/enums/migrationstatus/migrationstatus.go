package migrationstatus

type MigrationStatus struct {
	Name string
}

var (
	NOMIGRATIONNEEDED = MigrationStatus{"NO_MIGRATION_NEEDED"}

	MIGRATIONNEEDED = MigrationStatus{"MIGRATION_NEEDED"}

	INPROGRESS = MigrationStatus{"IN_PROGRESS"}

	ERROR = MigrationStatus{"ERROR"}
)

func (m MigrationStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MigrationStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m MigrationStatus) String() string { _ = "STUB: not implemented"; return "" }
