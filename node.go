package go_apache_config

import (
	"errors"
	"fmt"
)

type Node struct {
	Name     *string
	Content  *string
	Parent   *Node
	Children []*Node
}

// NewNode creates new root node.
func NewNode() *Node {
	return &Node{
		Name:     nil,
		Content:  nil,
		Parent:   nil,
		Children: nil,
	}
}

// CreateChildNode creates a child node and append a new child node into parent.
// A child node contains a configuration name and configuration content as well as a parent node in the tree.
// If the child node is an apache configuration section, it may also have child nodes of its own.
func (n *Node) CreateChildNode(name, content *string) (*Node, error) {
	if name == nil {
		return nil, errors.New("name should not be nil")
	}

	if content == nil {
		return nil, errors.New("content should not be nil")
	}

	child := &Node{
		Name:    name,
		Content: content,
		Parent:  n,
	}

	n.Children = append(n.Children, child)

	return child, nil
}

// DeleteChildNode deletes a node matched with name and content (which might be nil)
func (n *Node) DeleteChildNode(name, content *string) error {
	if name == nil {
		return errors.New("name should not be nil")
	}

	for i := 0; i < len(n.Children); i++ {
		if *n.Children[i].Name == *name &&
			(content == nil || (content != nil && *n.Children[i].Content == *content)) {
			copy(n.Children[i:], n.Children[i+1:])
			n.Children[len(n.Children)-1] = nil // release pointer reference to avoid memory leak
			n.Children = n.Children[:len(n.Children)-1]
			i--
		}
	}

	return nil
}

func (n *Node) String() string {
	var name, content string

	if n.Name != nil {
		name = *n.Name
	}
	if n.Content != nil {
		content = *n.Content
	}

	return fmt.Sprintf("Node {name=%s, content=%s, childNodeCount=%d}", name, content, len(n.Children))
}
