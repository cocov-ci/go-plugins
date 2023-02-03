package plugin

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/cocov-ci/go-plugin-kit/cocov"
	"github.com/cocov-ci/go-plugins/common"
	"go.uber.org/zap"
)

func Run(ctx cocov.Context) error {
	logger, err := common.SetupLogger("revive")
	if err != nil {
		log.Printf("Error configuring logger: %s", err)
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

	rootToml, rootTomlExists, err := reviveTomlExists(logger, ctx.Workdir())
	if err != nil {
		return nil, err
	}

	var issues []*common.CocovIssue
	for _, modPath := range modPaths {
		modDir := filepath.Dir(modPath)
		if err = common.GoModDownload(modDir, logger); err != nil {
			return nil, err
		}

		modToml, modTomlExists, err := reviveTomlExists(logger, modDir)
		if err != nil {
			return nil, err
		}

		var args []string
		switch {
		case modTomlExists:
			args = []string{"-formatter", "json", "--config", modToml}

		case !modTomlExists && rootTomlExists:
			args = []string{"-formatter", "json", "--config", rootToml}

		default:
			args = []string{"-formatter", "json"}
		}

		stdOut, stdErr, err := cocov.
			Exec2("revive", args, &cocov.ExecOpts{Workdir: modDir})

		if err != nil {
			logger.Error("Error running revive",
				zap.String("module path", modDir),
				zap.String("stdOut: ", string(stdOut)),
				zap.String("stdErr: ", string(stdErr)),
				zap.String("error:", err.Error()),
			)
			return nil, err
		}

		var modRules []rule
		if err := json.Unmarshal(stdOut, &modRules); err != nil {
			logger.Error("Error unmarshalling revive output",
				zap.String("output: ", string(stdOut)),
				zap.Error(err))
			return nil, err
		}

		for _, r := range modRules {
			kind, ok := common.ReviveRules[r.RuleName]
			if !ok {
				continue
			}

			filePath := filepath.Join(modDir, r.fileName())
			issue := common.NewCocovIssue(
				kind, r.startLine(), r.endLine(), filePath, r.text())

			issues = append(issues, issue)
		}
	}

	return issues, nil
}

func reviveTomlExists(logger *zap.Logger, path string) (string, bool, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		logger.Error(
			"Error searching for revive.toml",
			zap.String("path", path), zap.Error(err))
		return "", false, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if entry.Name() == "revive.toml" || entry.Name() == ".revive.toml" {
			return filepath.Join(path, entry.Name()), true, nil
		}
	}
	return "", false, nil
}
