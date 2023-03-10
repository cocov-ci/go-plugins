package plugin

import (
	"encoding/json"
	"os/exec"
	"path/filepath"

	"github.com/cocov-ci/go-plugin-kit/cocov"
	"github.com/cocov-ci/go-plugins/common"
	"go.uber.org/zap"
)

func Run(ctx cocov.Context) error {
	issues, err := run(ctx)
	if err != nil {
		return err
	}

	return common.EmitIssues(ctx, issues)
}

func run(ctx cocov.Context) ([]*common.CocovIssue, error) {
	modPaths, err := cocov.FindGoModules(ctx.Workdir())
	if err != nil {
		ctx.L().Error("Error searching for go modules",
			zap.String("path", ctx.Workdir()),
			zap.Error(err))
		return nil, err
	}

	var issues []*common.CocovIssue
	for _, modPath := range modPaths {
		modDir := filepath.Dir(modPath)
		opts := &cocov.ExecOpts{Workdir: modDir}

		stdOut, stdErr, err := cocov.Exec2("staticcheck", []string{"-f", "json", "./..."}, opts)
		if err != nil {
			if expectedErr, ok := err.(*exec.ExitError); !ok || expectedErr.ExitCode() != 1 {
				ctx.L().Error("Error running staticcheck",
					zap.String("path", modDir),
					zap.String("Std Err", string(stdErr)),
					zap.Error(err))
				return nil, err
			}
		}

		modIssues, err := parseChecks(stdOut, ctx.CommitSHA())
		if err != nil {
			ctx.L().Error("Error parsing checks",
				zap.String("module path", modDir),
				zap.Error(err))
			return nil, err
		}
		issues = append(issues, modIssues...)
	}

	return issues, nil
}

func parseChecks(stdOut []byte, commitSha string) ([]*common.CocovIssue, error) {
	var issues []*common.CocovIssue //nolint:prealloc
	var buff []byte
	for _, b := range stdOut {
		if b != '\n' {
			buff = append(buff, b)
			continue
		}

		c := check{}

		if err := json.Unmarshal(buff, &c); err != nil {
			return nil, err
		}

		buff = []byte{}

		kind, ok := staticChecks[c.Code]
		if !ok {
			continue
		}

		issue := common.NewCocovIssue(
			kind, c.Location.Line, c.End.Line,
			c.Location.File, c.Message, commitSha,
		)

		issues = append(issues, issue)
	}

	return issues, nil
}
