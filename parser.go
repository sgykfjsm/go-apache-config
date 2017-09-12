package go_apache_config

import (
	"bytes"
	"errors"
	"regexp"
)

var (
	// Assumed that the line will be trimmed space
	commentRegex      = regexp.MustCompile(`^#.*`)
	directiveRegex    = regexp.MustCompile(`([^\s]+)\s*(.+)`)
	sectionOpenRegex  = regexp.MustCompile(`^<([^/\s]+)\s*([^>]+)?>`)
	sectionCloseRegex = regexp.MustCompile(`^</([^\s]+)\s*>`)
)

func Parse(contents []byte) (*Node, error) {
	if contents == nil {
		return nil, errors.New("contents is empty")
	}

	var err error
	currentNode := NewNode()
	for _, line := range bytes.Split(contents, []byte("\n")) {
		line = bytes.TrimSpace(line)
		if len(line) == 0 || commentRegex.Match(line) {
			continue
		} else if sectionOpenRegex.Match(line) {
			group := sectionOpenRegex.FindSubmatch(line)
			if len(group) < 3 {
				continue
			}
			name := string(group[1])
			content := string(group[2])
			currentNode, err = currentNode.CreateChildNode(&name, &content)
			if err != nil {
				return nil, err
			}
		} else if sectionCloseRegex.Match(line) {
			currentNode = currentNode.Parent
		} else if directiveRegex.Match(line) {
			group := directiveRegex.FindSubmatch(line)
			if len(group) < 3 {
				continue
			}
			name := string(group[1])
			content := string(group[2])
			// CreateChildNode will return childNode
			_, err = currentNode.CreateChildNode(&name, &content)
			if err != nil {
				return nil, err
			}
		}
	}

	return currentNode, nil
}
