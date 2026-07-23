package indexprivilege

type IndexPrivilege struct {
	Name string
}

var (
	All = IndexPrivilege{"all"}

	Autoconfigure = IndexPrivilege{"auto_configure"}

	Create = IndexPrivilege{"create"}

	Createdoc = IndexPrivilege{"create_doc"}

	Createindex = IndexPrivilege{"create_index"}

	Createview = IndexPrivilege{"create_view"}

	Crossclusterreplication = IndexPrivilege{"cross_cluster_replication"}

	Crossclusterreplicationinternal = IndexPrivilege{"cross_cluster_replication_internal"}

	Delete = IndexPrivilege{"delete"}

	Deleteindex = IndexPrivilege{"delete_index"}

	Deleteview = IndexPrivilege{"delete_view"}

	Index = IndexPrivilege{"index"}

	Maintenance = IndexPrivilege{"maintenance"}

	Manage = IndexPrivilege{"manage"}

	Managedatastreamlifecycle = IndexPrivilege{"manage_data_stream_lifecycle"}

	Managefollowindex = IndexPrivilege{"manage_follow_index"}

	Manageilm = IndexPrivilege{"manage_ilm"}

	Manageleaderindex = IndexPrivilege{"manage_leader_index"}

	Manageview = IndexPrivilege{"manage_view"}

	Monitor = IndexPrivilege{"monitor"}

	None = IndexPrivilege{"none"}

	Read = IndexPrivilege{"read"}

	Readcrosscluster = IndexPrivilege{"read_cross_cluster"}

	Readviewmetadata = IndexPrivilege{"read_view_metadata"}

	Viewindexmetadata = IndexPrivilege{"view_index_metadata"}

	Write = IndexPrivilege{"write"}
)

func (i IndexPrivilege) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexPrivilege) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i IndexPrivilege) String() string { _ = "STUB: not implemented"; return "" }
