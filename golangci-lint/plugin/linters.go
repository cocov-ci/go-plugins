package plugin

import "github.com/cocov-ci/go-plugin-kit/cocov"

// descriptions about each linter can be found at:
// https://golangci-lint.run/usage/linters/
var linters = map[string]cocov.IssueKind{
	// Style
	"gosimple":          cocov.IssueKindStyle,
	"unused":            cocov.IssueKindStyle,
	"asciicheck":        cocov.IssueKindStyle,
	"containedctx":      cocov.IssueKindStyle,
	"deadcode":          cocov.IssueKindStyle,
	"decorder":          cocov.IssueKindStyle,
	"dogsled":           cocov.IssueKindStyle,
	"dupword":           cocov.IssueKindStyle,
	"errname":           cocov.IssueKindStyle,
	"exhaustivestruct":  cocov.IssueKindStyle,
	"exhaustruct":       cocov.IssueKindStyle,
	"forbidigo":         cocov.IssueKindStyle,
	"forcetypeassert":   cocov.IssueKindStyle,
	"gci":               cocov.IssueKindStyle,
	"ginkgolinter":      cocov.IssueKindStyle,
	"gochecknoglobals":  cocov.IssueKindStyle,
	"gochecknoinits":    cocov.IssueKindStyle,
	"goconst":           cocov.IssueKindStyle,
	"godot":             cocov.IssueKindStyle,
	"godox":             cocov.IssueKindStyle,
	"goerr113":          cocov.IssueKindStyle,
	"gofmt":             cocov.IssueKindStyle,
	"gofumpt":           cocov.IssueKindStyle,
	"goheader":          cocov.IssueKindStyle,
	"goimports":         cocov.IssueKindStyle,
	"golint":            cocov.IssueKindStyle,
	"gomnd":             cocov.IssueKindStyle,
	"gomoddirectives":   cocov.IssueKindStyle,
	"goprintffuncname":  cocov.IssueKindStyle,
	"grouper":           cocov.IssueKindStyle,
	"importas":          cocov.IssueKindStyle,
	"interfacebloat":    cocov.IssueKindStyle,
	"interfacer":        cocov.IssueKindStyle,
	"ireturn":           cocov.IssueKindStyle,
	"lll":               cocov.IssueKindStyle,
	"misspell":          cocov.IssueKindStyle,
	"nakedret":          cocov.IssueKindStyle,
	"nilnil":            cocov.IssueKindStyle,
	"nlreturn":          cocov.IssueKindStyle,
	"nolintlint":        cocov.IssueKindStyle,
	"nonamedreturns":    cocov.IssueKindStyle,
	"nosprintfhostport": cocov.IssueKindStyle,
	"predeclared":       cocov.IssueKindStyle,
	"promlinter":        cocov.IssueKindStyle,
	"structcheck":       cocov.IssueKindStyle,
	"stylecheck":        cocov.IssueKindStyle,
	"tagliatelle":       cocov.IssueKindStyle,
	"tenv":              cocov.IssueKindStyle,
	"thelper":           cocov.IssueKindStyle,
	"unconvert":         cocov.IssueKindStyle,
	"unparam":           cocov.IssueKindStyle,
	"usestdlibvars":     cocov.IssueKindStyle,
	"varcheck":          cocov.IssueKindStyle,
	"varnamelen":        cocov.IssueKindStyle,
	"wastedassign":      cocov.IssueKindStyle,
	"whitespace":        cocov.IssueKindStyle,
	"wrapcheck":         cocov.IssueKindStyle,
	"wsl":               cocov.IssueKindStyle,
	"structchek":        cocov.IssueKindStyle,

	// Convention
	"depguard":         cocov.IssueKindConvention,
	"gomodguard":       cocov.IssueKindConvention,
	"paralleltest":     cocov.IssueKindConvention,
	"testableexamples": cocov.IssueKindConvention,
	"testpackage":      cocov.IssueKindConvention,
	"ifshort":          cocov.IssueKindConvention,
	"nosnakecase":      cocov.IssueKindConvention,

	// Performance
	"maligned": cocov.IssueKindPerformance,
	"prealloc": cocov.IssueKindPerformance,

	// Security
	"gosec": cocov.IssueKindSecurity,

	// Bug
	"errcheck":      cocov.IssueKindBug,
	"govet":         cocov.IssueKindBug,
	"ineffassign":   cocov.IssueKindBug,
	"typecheck":     cocov.IssueKindBug,
	"asasalint":     cocov.IssueKindBug,
	"bidichk":       cocov.IssueKindBug,
	"bodyclose":     cocov.IssueKindBug,
	"contextcheck":  cocov.IssueKindBug,
	"durationcheck": cocov.IssueKindBug,
	"errchkjson":    cocov.IssueKindBug,
	"errorlint":     cocov.IssueKindBug,
	"execinquery":   cocov.IssueKindBug,
	"exhaustive":    cocov.IssueKindBug,
	"exportloopref": cocov.IssueKindBug,
	"loggercheck":   cocov.IssueKindBug,
	"makezero":      cocov.IssueKindBug,
	"nilerr":        cocov.IssueKindBug,
	"noctx":         cocov.IssueKindBug,
	"reassign":      cocov.IssueKindBug,
	"rowserrcheck":  cocov.IssueKindBug,
	"sqlclosecheck": cocov.IssueKindBug,
	"tparallel":     cocov.IssueKindBug,
	"scopelint":     cocov.IssueKindBug,

	// Complexity
	"cyclop":   cocov.IssueKindComplexity,
	"funlen":   cocov.IssueKindComplexity,
	"gocognit": cocov.IssueKindComplexity,
	"gocyclo":  cocov.IssueKindComplexity,
	"maintidx": cocov.IssueKindComplexity,
	"nestif":   cocov.IssueKindComplexity,

	// Duplication
	"dupl": cocov.IssueKindDuplication,

	// Exceptions
	"gocritic": issueKindException,
	"revive":   issueKindException,
}
