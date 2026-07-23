package bulk

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

func (r *Bulk) CreateOp(op types.CreateOperation, doc interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Bulk) IndexOp(op types.IndexOperation, doc interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Bulk) UpdateOp(op types.UpdateOperation, doc interface{}, update *types.UpdateAction) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Bulk) DeleteOp(op types.DeleteOperation) error { _ = "STUB: not implemented"; return nil }
