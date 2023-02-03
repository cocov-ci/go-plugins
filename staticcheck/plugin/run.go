package plugin

import (
	"encoding/json"
	"github.com/cocov-ci/go-plugin-kit/cocov"
	"github.com/cocov-ci/go-plugins/common"
	"go.uber.org/zap"
	"log"
	"os/exec"
	"path/filepath"
)

func Run(ctx cocov.Context) error {
	logger, err := common.SetupLogger("golangci-lint")
	if err != nil {
		log.Printf("Error configuring logger")
		return err
	}

	issues, err := run(ctx, logger)
	if err != nil {
		return err
	}

	return common.EmitIssues(ctx, logger, issues)
}

func run(ctx cocov.Context, logger *zap.Logger) ([]*common.CocovIssue, error) {
	modPaths, err := cocov.FindGoModules(ctx.Workdir())
	if err != nil {
		logger.Error("Error searching for go modules",
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
				logger.Error("Error running staticcheck",
					zap.String("path", modDir),
					zap.String("Std Err", string(stdErr)),
					zap.Error(err))
				return nil, err
			}
		}

		modIssues, err := parseChecks(stdOut)
		if err != nil {
			logger.Error("Error parsing checks",
				zap.String("module path", modDir),
				zap.Error(err))
			return nil, err
		}
		issues = append(issues, modIssues...)
	}

	return issues, nil
}

func parseChecks(stdOut []byte) ([]*common.CocovIssue, error) {
	var issues []*common.CocovIssue
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

		issue := common.NewCocovIssue(kind, c.Location.Line, c.End.Line, c.Location.File, c.Message)
		issues = append(issues, issue)
	}

	return issues, nil
}
