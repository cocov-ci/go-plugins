package common

import (
	"fmt"

	"github.com/cocov-ci/go-plugin-kit/cocov"
)

type CocovIssue struct {
	UID       string
	Kind      cocov.IssueKind
	FilePath  string
	LineStart uint
	LineEnd   uint
	Message   string
}

func NewCocovIssue(
	kind cocov.IssueKind,
	lineStart, lineEnd uint,
	filePath, message, commitSha string) *CocovIssue {
	c := &CocovIssue{
		Kind:      kind,
		FilePath:  filePath,
		LineStart: lineStart,
		LineEnd:   lineEnd,
		Message:   message,
	}
	c.hashID(commitSha)
	return c
}

func (c *CocovIssue) hashID(commitSha string) {
	input := fmt.Sprintf("%s-%s-%s-%s",
		commitSha, c.Kind, fmt.Sprintf("%d", c.LineStart), c.FilePath,
	)
	c.UID = cocov.SHA1([]byte(input))
}
