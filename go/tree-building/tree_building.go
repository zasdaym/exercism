//Package tree is solution for problem Tree Building.
package tree

import (
	"fmt"
	"sort"
)

const rootID = 0

type (
	// Node represents a node in a tree.
	Node struct {
		ID       int
		Children []*Node
	}

	// Record represents a post record.
	Record struct {
		ID     int
		Parent int
	}
)

var (
	// ErrInvalidRecordID represents an invalid Record ID
	ErrInvalidRecordID = fmt.Errorf("invalid record ID, must be between 0 - records length")

	// ErrNonContinuousNode represents non-continuous input records ID
	ErrNonContinuousNode = fmt.Errorf("non-continuous node")

	// ErrInvalidParent represents invalid Parent ID on a Record.
	ErrInvalidParent = fmt.Errorf("node has invalid parent")
)

// Build creates a Tree built from given Records and returns its root node.
func Build(records []Record) (*Node, error) {
	// edge case for zero record
	if len(records) == 0 {
		return nil, nil
	}

	// sort given records by ID
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// range over sorted records
	nodes := make([]Node, len(records))
	for i, r := range records {
		// check if record ID doesn't match the sorted index
		if r.ID != i {
			return nil, ErrNonContinuousNode
		}

		// check if parent ID of a record is valid
		isValidParentID := (r.ID > r.Parent) || (r.ID == rootID && r.Parent == rootID)
		if !isValidParentID {
			return nil, ErrInvalidParent
		}

		// append current record to tree
		nodes[i].ID = i
		if i != rootID {
			// add current record as children of its parent
			parent := &nodes[r.Parent]
			parent.Children = append(parent.Children, &nodes[i])
		}
	}

	return &nodes[0], nil
}
