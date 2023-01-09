package plugin

import "github.com/cocov-ci/go-plugin-kit/cocov"

// issueKindException is used only for identifying a kind at linters
// and must not be used as a real IssueKind.
const issueKindException cocov.IssueKind = 255

const (
	exceptionGoCritic = "gocritic"
	exceptionRevive   = "revive"
)
