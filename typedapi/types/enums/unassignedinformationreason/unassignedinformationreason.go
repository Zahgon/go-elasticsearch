package unassignedinformationreason

type UnassignedInformationReason struct {
	Name string
}

var (
	INDEXCREATED = UnassignedInformationReason{"INDEX_CREATED"}

	CLUSTERRECOVERED = UnassignedInformationReason{"CLUSTER_RECOVERED"}

	INDEXREOPENED = UnassignedInformationReason{"INDEX_REOPENED"}

	DANGLINGINDEXIMPORTED = UnassignedInformationReason{"DANGLING_INDEX_IMPORTED"}

	NEWINDEXRESTORED = UnassignedInformationReason{"NEW_INDEX_RESTORED"}

	EXISTINGINDEXRESTORED = UnassignedInformationReason{"EXISTING_INDEX_RESTORED"}

	REPLICAADDED = UnassignedInformationReason{"REPLICA_ADDED"}

	ALLOCATIONFAILED = UnassignedInformationReason{"ALLOCATION_FAILED"}

	NODELEFT = UnassignedInformationReason{"NODE_LEFT"}

	REROUTECANCELLED = UnassignedInformationReason{"REROUTE_CANCELLED"}

	REINITIALIZED = UnassignedInformationReason{"REINITIALIZED"}

	REALLOCATEDREPLICA = UnassignedInformationReason{"REALLOCATED_REPLICA"}

	PRIMARYFAILED = UnassignedInformationReason{"PRIMARY_FAILED"}

	FORCEDEMPTYPRIMARY = UnassignedInformationReason{"FORCED_EMPTY_PRIMARY"}

	MANUALALLOCATION = UnassignedInformationReason{"MANUAL_ALLOCATION"}
)

func (u UnassignedInformationReason) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *UnassignedInformationReason) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (u UnassignedInformationReason) String() string { _ = "STUB: not implemented"; return "" }
