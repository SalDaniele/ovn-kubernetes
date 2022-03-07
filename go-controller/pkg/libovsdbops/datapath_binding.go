package libovsdbops

import (
	libovsdbclient "github.com/ovn-org/libovsdb/client"

	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/sbdb"
)

type datapathBindingPredicate func(*sbdb.DatapathBinding) bool

// GetDatapathBindingWithPredicate looks up a datapath binding from the cache
// based on a given predicate. If none or multiple are found, an error is
// returned.
func GetDatapathBindingWithPredicate(sbClient libovsdbclient.Client, p datapathBindingPredicate) (*sbdb.DatapathBinding, error) {
	found := []*sbdb.DatapathBinding{}
	opModel := OperationModel{
		ModelPredicate: p,
		ExistingResult: &found,
		OnModelUpdates: nil, // no update
		ErrNotFound:    true,
		BulkOp:         false,
	}

	m := NewModelClient(sbClient)
	_, err := m.CreateOrUpdate(opModel)
	if err != nil {
		return nil, err
	}

	return found[0], nil
}
