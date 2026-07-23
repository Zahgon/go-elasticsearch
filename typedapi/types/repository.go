package types

type Repository any

type RepositoryVariant interface {
	RepositoryCaster() *Repository
}
