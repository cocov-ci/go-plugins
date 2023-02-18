package plugin

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocov-ci/go-plugins/common"
	sdkmocks "github.com/cocov-ci/go-plugins/common/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// TestRun assumes that the runner has golangci-lint installed.
func TestRun(t *testing.T) {
	root := common.FindParentDir(t)
	l := zap.NewNop()

	t.Run("Works as expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)

		fixtureYamlPath := filepath.
			Join(root, "golangci-lint", "fixtures", "fixture-file.yaml")

		data, err := os.ReadFile(fixtureYamlPath)
		require.NoError(t, err)

		targetFilePath := filepath.
			Join(root, "golangci-lint", ".golangci.yaml")

		file, err := os.OpenFile(targetFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
		require.NoError(t, err)

		_, err = file.Write(data)
		require.NoError(t, err)

		defer file.Close()
		defer os.Remove(targetFilePath)

		ctx.EXPECT().Workdir().Return(root)

		issues, err := run(ctx, l)
		require.NoError(t, err)

		for _, issue := range issues {
			info, err := os.Stat(issue.FilePath)
			require.NoError(t, err)
			require.False(t, info.IsDir())
		}
	})
}
