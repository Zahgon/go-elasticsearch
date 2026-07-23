package query

import (
	"context"
	"encoding/json"
)

type metadata struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type esqlResponse struct {
	Columns []metadata `json:"columns"`
	Values  [][]any    `json:"values"`
}

func Helper[T any](ctx context.Context, esqlQuery *Query) ([]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type EsqlIterator[T any] interface {
	Next() (*T, error)
	More() bool
}

type iterator[T any] struct {
	reader    []byte
	decoder   *json.Decoder
	keys      []string
	skipComma bool
}

func (d iterator[T]) More() bool { _ = "STUB: not implemented"; return false }

func (d iterator[T]) Next() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIteratorHelper[T any](ctx context.Context, query *Query) (EsqlIterator[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
