package go_apache_config

import (
	"fmt"
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode()
	if node.Name != nil || node.Content != nil ||
		node.Parent != nil || node.Children != nil {
		t.Error("Each properties should be nil")
	}
}

func TestNode_CreateChildNode_WithNilName_ReturnError(t *testing.T) {
	root := NewNode()
	content := "testContent"

	_, err := root.CreateChildNode(nil, &content)
	if err == nil {
		t.Error("when name is nil, CreateChildNode should return error")
	}
}

func TestNode_CreateChildNode_WithNilContent_ReturnError(t *testing.T) {
	root := NewNode()
	name := "testNode"

	_, err := root.CreateChildNode(&name, nil)
	if err == nil {
		t.Error("when content is nil, CreateChildNode should return error")
	}
}

func TestNode_CreateChildNode(t *testing.T) {
	root := NewNode()
	name := "testNode"
	content := "testContent"

	before := len(root.Children)
	_, err := root.CreateChildNode(&name, &content)
	if err != nil {
		t.Error(err.Error())
	}

	after := len(root.Children)
	if after-1 != before {
		t.Error("After root created child, root's children should be increased.")
	}
}

func TestNode_DeleteChildNode_WithNilName_ReturnError(t *testing.T) {
	root := NewNode()
	name := "testNode"
	content := "testContent"

	if _, err := root.CreateChildNode(&name, &content); err != nil {
		t.Error(err.Error())
	}

	if err := root.DeleteChildNode(nil, &content); err == nil {
		t.Error("when name is nil, DeleteChildNode should return error")
	}
}

func TestNode_DeleteChildNode_WithNilName_Works(t *testing.T) {
	root := NewNode()
	name := "testNode"
	content := "testContent"

	if _, err := root.CreateChildNode(&name, &content); err != nil {
		t.Error(err.Error())
	}

	if err := root.DeleteChildNode(&name, nil); err != nil {
		t.Error(err.Error())
	}
}

func TestNode_DeleteChildNode(t *testing.T) {
	root := NewNode()
	name := "testNode"
	content := "testContent"

	if _, err := root.CreateChildNode(&name, &content); err != nil {
		t.Error(err.Error())
	}

	name2 := "testNode2"
	content2 := "testContent2"
	if _, err := root.CreateChildNode(&name2, &content2); err != nil {
		t.Error(err.Error())
	}

	before := len(root.Children)
	if err := root.DeleteChildNode(&name, &content); err != nil {
		t.Error(err.Error())
	}

	after := len(root.Children)
	if after != 1 && after+1 == before {
		t.Error("After childnode has been deleted, there should be no children in the root.")
	}

	if *root.Children[0].Name != name2 {
		t.Error(fmt.Sprintf("expected %q but actual %q", name2, *root.Children[0].Name))
	}

	if *root.Children[0].Content != content2 {
		t.Error(fmt.Sprintf("expected %q but actual %q", content2, *root.Children[0].Content))
	}
}

func TestNode_String(t *testing.T) {
	root := NewNode()

	childrenNum := 3
	for i := 0; i < childrenNum; i++ {
		name := fmt.Sprintf("testNode%d", i)
		content := fmt.Sprintf("testContent%d", i)
		if _, err := root.CreateChildNode(&name, &content); err != nil {
			t.Error(err.Error())
		}
	}

	expect := fmt.Sprintf("Node {name=, content=, childNodeCount=%d}", childrenNum)
	actual := root.String()
	if expect != actual {
		t.Error(fmt.Sprintf("root.String() should return %q but %q", expect, actual))
	}

	child := root.Children[0]
	expect = fmt.Sprint("Node {name=testNode0, content=testContent0, childNodeCount=0}")
	actual = child.String()
	if expect != actual {
		t.Error(fmt.Sprintf("root.String() should return %q but %q", expect, actual))
	}

}
