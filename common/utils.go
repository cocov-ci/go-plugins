package common

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocov-ci/go-plugin-kit/cocov"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// EmitIssues emits one report for each provided issue. Logs and returns
// the error if it fails.
func EmitIssues(ctx cocov.Context, log *zap.Logger, issues []*CocovIssue) error {
	for _, issue := range issues {
		err := ctx.EmitIssue(
			issue.Kind, issue.FilePath, issue.LineStart,
			issue.LineEnd, issue.Message, issue.UID)

		if err != nil {
			log.Error("Error issuing cocov report", zap.Error(err))
			return err
		}
	}
	return nil
}

// SetupLogger configures a logger based on COCOV_ENV environment variable and
// returns a logger named with the pluginName or an error if it fails.
func SetupLogger(pluginName string) (*zap.Logger, error) {
	l, err := zap.NewProduction()
	if os.Getenv("COCOV_ENV") == "development" {
		opts := zap.Development()
		l, err = zap.NewDevelopment(opts)
	}
	if err != nil {
		return nil, err
	}
	return l.With(zap.String("plugin", pluginName)), nil
}

// GoModDownload runs the command `go mod download` at a given path.
// Logs and returns the error if it fails.
func GoModDownload(path string, log *zap.Logger) error {
	args := []string{"mod", "download"}
	out, err := cocov.Exec("go", args, &cocov.ExecOpts{Workdir: path})

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
func FindParentDir(t *testing.T, actual, target string) string {
	require.NotEqual(t, actual, "")

	if filepath.Base(actual) == target {
		return actual
	}

	return FindParentDir(t, filepath.Dir(actual), target)
}
