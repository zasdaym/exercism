//Package tree is solution for problem Tree Building.
package tree

import (
	"fmt"
)

const rootID = 0

type (
	Node struct {
		ID int
		Children []*Node
	}

	Record struct {
		ID int
		Parent int
	}
)

var ErrInvalidRecordID = fmt.Errorf("invalid record ID, must be between 0 - records length")
var ErrNonContinuousNode = fmt.Errorf("non-continuous node")
var ErrInvalidParent = fmt.Errorf("node has invalid parent")

// Build returns a root node of built tree from given records.
func Build(records []Record) (*Node, error) {
	// edge case for zero record
	if len(records) == 0 {
		return nil, nil
	}

	// sort given node records by its ID
	positions := make([]int, len(records))
	for i, r := range records {
		if r.ID < rootID || r.ID >= len(records) {
			return nil, ErrInvalidRecordID
		}

		positions[r.ID] = i
	}

	// range over sorted node records
	nodes := make([]Node, len(records))
	for i := range positions {
		r := records[positions[i]]
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
			// add current record as children of parent node
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}

	return &nodes[0], nil
}
