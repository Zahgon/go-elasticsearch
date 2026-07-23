package types

type IndicesPrivilegesQuery any

type IndicesPrivilegesQueryVariant interface {
	IndicesPrivilegesQueryCaster() *IndicesPrivilegesQuery
}
