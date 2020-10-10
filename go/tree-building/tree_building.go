//Package tree is solution for problem Tree Building.
package tree

import (
	"fmt"
	"sort"
)

const rootID = 0

// Node represents a node in a tree.
type Node struct {
	ID       int
	Children []*Node
}

// Record represents a post record.
type Record struct {
	ID     int
	Parent int
}

// Build creates a Tree built from given Records and returns its root node.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	nodes := make([]Node, len(records))
	for i, r := range records {
		if r.ID != i || r.ID < r.Parent || r.ID != rootID && r.ID == r.Parent {
			return nil, fmt.Errorf("invalid record: %v", r)
		}

		nodes[i].ID = i
		if i != rootID {
			parent := &nodes[r.Parent]
			parent.Children = append(parent.Children, &nodes[i])
		}
	}
	return &nodes[0], nil
}
