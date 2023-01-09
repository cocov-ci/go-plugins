package plugin

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocov-ci/go-plugins/common"
	sdkmocks "github.com/cocov-ci/go-plugins/common/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestRun(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	parent := common.FindParentDir(t, wd, "go-plugins")
	l := zap.NewNop()
	tomlFileName := "revive.toml"

	t.Run("fails with malformed toml", func(t *testing.T) {
		tomlPath := filepath.Join(parent, "revive", "fixtures", "malformed.toml")

		data, err := os.ReadFile(tomlPath)
		require.NoError(t, err)

		localTomlPath := filepath.Join(parent, "revive", "revive.toml")
		f, err := os.OpenFile(localTomlPath, os.O_RDWR|os.O_CREATE, 0666)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write(data)
		require.NoError(t, err)

		defer os.Remove(localTomlPath)

		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)
		ctx.EXPECT().Workdir().Return(parent).MaxTimes(2)

		_, err = run(ctx, l)
		assert.Error(t, err)
	})

	t.Run("without toml", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)
		ctx.EXPECT().Workdir().Return(parent).MaxTimes(2)

		issues, err := run(ctx, l)
		require.NoError(t, err)

		for _, issue := range issues {
			info, err := os.Stat(issue.FilePath)
			require.NoError(t, err)
			assert.False(t, info.IsDir())
		}
	})

	t.Run("with root toml", func(t *testing.T) {
		tomlPath := filepath.Join(parent, "revive", "fixtures", tomlFileName)

		data, err := os.ReadFile(tomlPath)
		require.NoError(t, err)

		parentTomlPath := filepath.Join(parent, "revive.toml")
		f, err := os.OpenFile(parentTomlPath, os.O_RDWR|os.O_CREATE, 0666)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write(data)
		require.NoError(t, err)

		defer os.Remove(parentTomlPath)

		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)
		ctx.EXPECT().Workdir().Return(parent).MaxTimes(2)

		issues, err := run(ctx, l)
		assert.NoError(t, err)

		for _, issue := range issues {
			info, err := os.Stat(issue.FilePath)
			assert.NoError(t, err)
			assert.False(t, info.IsDir())
		}
	})

	t.Run("with local toml", func(t *testing.T) {
		tomlPath := filepath.Join(parent, "revive", "fixtures", tomlFileName)

		data, err := os.ReadFile(tomlPath)
		require.NoError(t, err)

		localTomlPath := filepath.Join(parent, "revive", "revive.toml")
		f, err := os.OpenFile(localTomlPath, os.O_RDWR|os.O_CREATE, 0666)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write(data)
		require.NoError(t, err)

		defer os.Remove(localTomlPath)

		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)
		ctx.EXPECT().Workdir().Return(parent).MaxTimes(2)

		issues, err := run(ctx, l)
		assert.NoError(t, err)

		for _, issue := range issues {
			info, err := os.Stat(issue.FilePath)
			assert.NoError(t, err)
			assert.False(t, info.IsDir())
		}
	})
}

func TestReviveTomlExists(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	parent := common.FindParentDir(t, wd, "go-plugins")
	fixtureTomlPath := filepath.Join(parent, "revive/fixtures")
	_, ok, err := reviveTomlExists(zap.NewNop(), fixtureTomlPath)
	require.NoError(t, err)
	assert.True(t, ok)
}
