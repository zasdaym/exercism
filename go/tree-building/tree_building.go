// Package tree provides Tree Building solution.
package tree

import (
	"errors"
	"fmt"
	"sort"
)

// Record represents a tree record.
type Record struct {
	ID     int
	Parent int
}

// Node represents a tree node.
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a new tree from given records.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node)
	lastID := -1
	for _, record := range records {
		if record.ID != lastID+1 {
			return nil, fmt.Errorf("non-continuous node ID %d after node ID %d", record.ID, lastID)
		}

		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("root node has parent")
		}

		if record.ID != 0 && record.ID == record.Parent {
			return nil, fmt.Errorf("tree cycle directly on node with ID %d", record.ID)
		}

		lastID = record.ID

		node := &Node{
			ID: record.ID,
		}
		nodes[record.ID] = node

		// Do not append child of the first record.
		if record.ID == 0 {
			continue
		}

		parent, exists := nodes[record.Parent]
		if !exists {
			return nil, fmt.Errorf("parent node with ID %d doesn't exists", record.Parent)
		}
		parent.Children = append(parent.Children, node)
	}

	return nodes[0], nil
}
