package go_apache_config

import (
	"io/ioutil"
	"log"
	"testing"
)

const resourceConf = "resource/test/resource.conf"

func readResourceConf() []byte {
	content, err := ioutil.ReadFile(resourceConf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return content
}

func verifyNode(t *testing.T, node *Node, name, content *string, childrenNum int) error {
	if name == nil {
		if node.Name != nil {
			t.Errorf("node.Name: expected nil but actual %v", *node.Name)
		}
	} else { // name != nil
		if node.Name == nil {
			t.Errorf("node.Name: expected %v but actual nil", name)
		} else if *node.Name != *name {
			t.Errorf("node.Name: expected %v but actual %v", name, *node.Name)
		}
	}

	if content == nil {
		if node.Content != nil {
			t.Errorf("node.Content: expected nil but actual %v", *node.Content)
		}
	} else { // content != nil
		if node.Content == nil {
			t.Errorf("node.Content: expected %v but actual nil", content)
		} else if *node.Content != *content {
			t.Errorf("node.Content: expected %v but actual %v", content, *node.Content)
		}
	}

	if len(node.Children) != childrenNum {
		t.Errorf("node.Children: expected children count is %d but actual is %d", childrenNum, len(node.Children))
	}

	return nil
}

func TestParse_WithNilContent_ReturnError(t *testing.T) {
	if _, err := Parse(nil); err == nil {
		t.Error("when contents is nil, Parse should return error")
	}
}

func TestParse(t *testing.T) {
	contents := readResourceConf()
	node, err := Parse(contents)
	if err != nil {
		t.Error(err.Error())
	}

	if err := verifyNode(t, node, nil, nil, 2); err != nil {
		t.Error(err.Error())
	}
}
