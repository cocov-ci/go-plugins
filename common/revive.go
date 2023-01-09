package common

import "github.com/cocov-ci/go-plugin-kit/cocov"

// ReviveRules represents the set of rules used by revive.
// All revive rules can be found at:
// https://github.com/mgechev/revive/blob/master/RULES_DESCRIPTIONS.md#add-constant
var ReviveRules = map[string]cocov.IssueKind{
	// Style
	"confusing-naming":  cocov.IssueKindStyle,
	"confusing-results": cocov.IssueKindStyle,
	"early-return":      cocov.IssueKindStyle,
	"empty-lines":       cocov.IssueKindStyle,

	// Convention
	"argument-limit":        cocov.IssueKindConvention,
	"banned-characters":     cocov.IssueKindConvention,
	"context-as-argument":   cocov.IssueKindConvention,
	"error-naming":          cocov.IssueKindConvention,
	"function-result-limit": cocov.IssueKindConvention,
	"increment-decrement":   cocov.IssueKindConvention,
	"imports-blacklist":     cocov.IssueKindConvention,
	"line-length-limit":     cocov.IssueKindConvention,
	"max-public-structs":    cocov.IssueKindConvention,
	"nested-structs":        cocov.IssueKindConvention,
	"package-comments":      cocov.IssueKindConvention,
	"receiver-naming":       cocov.IssueKindConvention,
	"time-naming":           cocov.IssueKindConvention,
	"var-naming":            cocov.IssueKindConvention,
	"var-declaration":       cocov.IssueKindConvention,
	"file-header":           cocov.IssueKindConvention,
	"modifies-parameter":    cocov.IssueKindConvention,

	// Quality
	"add-constant":          cocov.IssueKindQuality,
	"blank-imports":         cocov.IssueKindQuality,
	"bool-literal-in-expr":  cocov.IssueKindQuality,
	"constant-logical-expr": cocov.IssueKindQuality,
	"dot-imports":           cocov.IssueKindQuality,
	"empty-block":           cocov.IssueKindQuality,
	"error-return":          cocov.IssueKindQuality,
	"exported":              cocov.IssueKindQuality,
	"get-return":            cocov.IssueKindQuality,
	"identical-branches":    cocov.IssueKindQuality,
	"indent-error-flow":     cocov.IssueKindQuality,
	"import-shadowing":      cocov.IssueKindQuality,
	"unnecessary-stmt":      cocov.IssueKindQuality,
	"unreachable-code":      cocov.IssueKindQuality,
	"unused-parameter":      cocov.IssueKindQuality,
	"unused-receiver":       cocov.IssueKindQuality,
	"useless-break":         cocov.IssueKindQuality,
	"string-format":         cocov.IssueKindQuality,
	"superfluous-else":      cocov.IssueKindQuality,
	"unexported-naming":     cocov.IssueKindQuality,
	"unexported-return":     cocov.IssueKindQuality,
	"flag-parameter":        cocov.IssueKindQuality,
	"if-return":             cocov.IssueKindQuality,

	// Performance
	"optimize-operands-order": cocov.IssueKindPerformance,

	// Bug
	"atomic":                  cocov.IssueKindBug,
	"redefines-builtin-id":    cocov.IssueKindBug,
	"string-of-int":           cocov.IssueKindBug,
	"struct-tag":              cocov.IssueKindBug,
	"time-equal":              cocov.IssueKindBug,
	"unconditional-recursion": cocov.IssueKindBug,
	"unhandled-error":         cocov.IssueKindBug,
	"waitgroup-by-value":      cocov.IssueKindBug,
	"bare-return":             cocov.IssueKindBug,
	"deep-exit":               cocov.IssueKindBug,
	"modifies-value-receiver": cocov.IssueKindBug,
	"range-val-in-closure":    cocov.IssueKindBug,
	"call-to-gc":              cocov.IssueKindBug,
	"context-keys-type":       cocov.IssueKindBug,
	"datarace":                cocov.IssueKindBug,
	"defer":                   cocov.IssueKindBug,

	// Complexity
	"cognitive-complexity": cocov.IssueKindComplexity,
	"cyclomatic":           cocov.IssueKindComplexity,
	"function-length":      cocov.IssueKindComplexity,

	// Duplication
	"duplicated-imports": cocov.IssueKindDuplication,
}
