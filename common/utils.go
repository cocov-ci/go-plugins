package common

import (
	"strings"
	"testing"

	"github.com/cocov-ci/go-plugin-kit/cocov"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// EmitIssues emits one report for each provided issue. Logs and returns
// the error if it fails.
func EmitIssues(ctx cocov.Context, issues []*CocovIssue) error {
	for _, issue := range issues {
		err := ctx.EmitIssue(
			issue.Kind, issue.FilePath, issue.LineStart,
			issue.LineEnd, issue.Message, issue.UID)

		if err != nil {
			ctx.Logger().Error("Error issuing cocov report", zap.Error(err))
			return err
		}
	}
	return nil
}

// GoModDownload runs the command `go mod download` at a given path.
// Logs and returns the error if it fails.
func GoModDownload(path string, log *zap.Logger, envs map[string]string) error {
	args := []string{"mod", "download"}
	out, err := cocov.Exec("go", args, &cocov.ExecOpts{Workdir: path, Env: envs})

	if err != nil {
		log.Error(
			"Error downloading go modules",
			zap.Error(err), zap.String("output:", string(out)))
		return err
	}

	return nil
}

// FindParentDir is a facility function used for testing
// that returns the path to a parent directory based on the actual path.
// It also stops the execution if the target path can not be found.
func FindParentDir(t *testing.T) string {
	out, err := cocov.Exec("git", []string{"rev-parse", "--show-toplevel"}, nil)
	require.NoError(t, err)
	return strings.TrimSpace(string(out))
}
