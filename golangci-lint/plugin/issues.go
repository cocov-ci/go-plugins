package plugin

import (
	"path/filepath"
	"strings"

	"github.com/cocov-ci/go-plugin-kit/cocov"
	"github.com/cocov-ci/go-plugins/common"
)

type goCILintOutput struct {
	Issues []Issue `json:"Issues"`
}

type Issue struct {
	FromLinter string `json:"FromLinter"`
	Text       string `json:"Text"`
	Pos        Pos    `json:"Pos"`
}

type Pos struct {
	Filename string `json:"Filename"`
	Line     uint   `json:"Line"`
}

func newCocovIssue(rootPath, commitSha string, i Issue, kind cocov.IssueKind) (*common.CocovIssue, bool) {
	if kind == issueKindException {
		return newExceptionIssue(rootPath, commitSha, i)
	}

	filePath := filepath.Join(rootPath, i.Pos.Filename)
	ccIssue := common.NewCocovIssue(kind, i.Pos.Line, i.Pos.Line, filePath, i.Text, commitSha)
	return ccIssue, true
}

func newExceptionIssue(rootPath, commitSha string, i Issue) (*common.CocovIssue, bool) {
	// go-critic and revive report messages are presented
	// in the following format:
	// checker:message
	splitMessage := strings.Split(i.Text, ":")
	if len(splitMessage) < 2 {
		return nil, false
	}

	identifier := splitMessage[0]

	var ok bool
	var kind cocov.IssueKind

	switch i.FromLinter {
	case exceptionGoCritic:
		kind, ok = common.GoCriticCheckers[identifier]
	case exceptionRevive:
		kind, ok = common.ReviveRules[identifier]
	}

	if !ok {
		return nil, false
	}

	text := splitMessage[1]
	filePath := filepath.Join(rootPath, i.Pos.Filename)
	issue := common.
		NewCocovIssue(kind, i.Pos.Line, i.Pos.Line, filePath, text, commitSha)

	return issue, true
}
