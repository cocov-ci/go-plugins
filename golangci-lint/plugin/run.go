package plugin

import (
	"encoding/json"
	"os"
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
		ctx.L().Error(
			"Error searching for go modules",
			zap.String("path", ctx.Workdir()),
			zap.Error(err),
		)
		return nil, err
	}

	var repoIssues []*common.CocovIssue
	for _, modPath := range modPaths {
		modDir := filepath.Dir(modPath)
		sumPath := filepath.Join(modDir, "go.sum")

		cachePath, err := os.MkdirTemp("", "")
		if err != nil {
			panic(err)
		}

		goEnv := map[string]string{
			"GOMODCACHE": cachePath,
		}

		keys := []string{modPath, sumPath}
		if _, err = ctx.LoadArtifactCache(keys, cachePath); err != nil {
			ctx.L().Error("Error loading cache artifact", zap.Error(err))
			return nil, err
		}

		ctx.L().Info("Working", zap.String("at", modDir))

		if err = common.GoModDownload(modDir, ctx.L(), goEnv); err != nil {
			return nil, err
		}

		// bad file
		if err = ctx.StoreArtifactCache(keys, cachePath); err != nil {
			ctx.L().Error("Error storing cache artifact", zap.Error(err))
			return nil, err
		}

		output, err := runGolangCILint(modDir, ctx.L(), goEnv)
		if err != nil {
			return nil, err
		}

		if output == nil {
			continue
		}

		modIssues := buildCocovIssues(modDir, ctx.CommitSHA(), output)
		repoIssues = append(repoIssues, modIssues...)
	}
	return repoIssues, nil
}

func runGolangCILint(path string, log *zap.Logger, env map[string]string) (*goCILintOutput, error) {
	args := []string{"run", "--out-format", "json"}
	stdOut, stdErr, err := cocov.
		Exec2("golangci-lint", args, &cocov.ExecOpts{Workdir: path, Env: env})

	if err != nil {
		if !isGolangCILintExpectedErr(err) {
			log.Error("Error executing golangci-lint", zap.Error(err),
				zap.String("Std error: ", string(stdErr)))
			return nil, err
		}
	}

	if len(stdOut) < 2 {
		return nil, nil
	}

	out := &goCILintOutput{}
	if err = json.Unmarshal(stdOut, out); err != nil {
		log.Error("Error unmarshalling output",
			zap.Error(err),
			zap.String("output:", string(stdOut)))
		return nil, err
	}

	if len(out.Issues) == 0 {
		return nil, nil
	}

	return out, nil
}

func buildCocovIssues(path, commitSha string, out *goCILintOutput) []*common.CocovIssue {
	ccIssues := make([]*common.CocovIssue, 0, len(out.Issues))
	for _, issue := range out.Issues {
		kind, ok := linters[issue.FromLinter]
		if !ok {
			continue
		}

		i, ok := newCocovIssue(path, commitSha, issue, kind)
		if !ok {
			continue
		}
		ccIssues = append(ccIssues, i)
	}
	return ccIssues
}

func isGolangCILintExpectedErr(err error) bool {
	extErr, ok := err.(*exec.ExitError)
	return ok && extErr.ExitCode() == 1
}
