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
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	if err := checkRecords(records); err != nil {
		return nil, err
	}

	nodes := make([]Node, len(records))
	for i, r := range records {
		nodes[i].ID = i
		if i != rootID {
			parent := &nodes[r.Parent]
			parent.Children = append(parent.Children, &nodes[i])
		}
	}

	return &nodes[0], nil
}

func checkRecords(records []Record) error {
	for i, r := range records {
		if r.ID != i {
			return ErrNonContinuousNode
		}

		isValidParentID := (r.ID > r.Parent) || (r.ID == rootID && r.Parent == rootID)
		if !isValidParentID {
			return ErrInvalidParent
		}
	}
	return nil
}
