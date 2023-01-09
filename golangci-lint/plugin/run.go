package plugin

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/cocov-ci/go-plugin-kit/cocov"
	"github.com/cocov-ci/go-plugins/common"
	"go.uber.org/zap"
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
		logger.Error(
			"Error searching for go modules",
			zap.String("path", ctx.Workdir()),
			zap.Error(err),
		)
		return nil, err
	}

	var repoIssues []*common.CocovIssue
	for _, modPath := range modPaths {
		modDir := filepath.Dir(modPath)
		logger.Info("Working", zap.String("at", modDir))

		if err = common.GoModDownload(modDir, logger); err != nil {
			return nil, err
		}

		output, err := runGolangCILint(modDir, logger)
		if err != nil {
			return nil, err
		}

		modIssues := buildCocovIssues(modDir, output)
		repoIssues = append(repoIssues, modIssues...)
	}
	return repoIssues, nil
}

func runGolangCILint(path string, log *zap.Logger) (*goCILintOutput, error) {
	args := []string{"run", "--out-format", "json"}
	stdOut, stdErr, err := cocov.
		Exec2("golangci-lint", args, &cocov.ExecOpts{Workdir: path})

	if err != nil {
		if !isGolangCILintExpectedErr(err) {
			log.Error("Error executing golangci-lint", zap.Error(err),
				zap.String("Std error: ", string(stdErr)))
			return nil, err
		}

	} else if len(stdOut) == 0 {
		log.Error("Std out is empty", zap.String("Std error:", string(stdErr)))
		return nil, fmt.Errorf("std error :%s", string(stdErr))
	}

	out := &goCILintOutput{}
	if err = json.Unmarshal(stdOut, out); err != nil {
		log.Error("Error unmarshalling output",
			zap.Error(err),
			zap.String("output:", string(stdOut)))
		return nil, err
	}

	return out, nil
}

func buildCocovIssues(path string, out *goCILintOutput) []*common.CocovIssue {
	ccIssues := make([]*common.CocovIssue, 0, len(out.Issues))
	for _, issue := range out.Issues {
		kind, ok := linters[issue.FromLinter]
		if !ok {
			continue
		}

		i, ok := newCocovIssue(path, issue, kind)
		if !ok {
			continue
		}
		ccIssues = append(ccIssues, i)
	}
	return ccIssues
}

func isGolangCILintExpectedErr(err error) bool {
	extErr, ok := err.(*exec.ExitError)
	return ok || extErr.ExitCode() != 1
}
