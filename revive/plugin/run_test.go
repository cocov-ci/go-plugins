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
	root := common.FindParentDir(t)
	l := zap.NewNop()
	tomlFileName := "revive.toml"
	sha := "sha"

	t.Run("fails with malformed toml", func(t *testing.T) {
		tomlPath := filepath.Join(root, "revive", "fixtures", "malformed.toml")

		data, err := os.ReadFile(tomlPath)
		require.NoError(t, err)

		localTomlPath := filepath.Join(root, "revive", "revive.toml")
		f, err := os.OpenFile(localTomlPath, os.O_RDWR|os.O_CREATE, 0666)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write(data)
		require.NoError(t, err)

		defer os.Remove(localTomlPath)

		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)
		ctx.EXPECT().Workdir().Return(root).MaxTimes(2)
		ctx.EXPECT().L().
			DoAndReturn(func() *zap.Logger { return l }).
			AnyTimes()
		ctx.EXPECT().CommitSHA().Return(sha).AnyTimes()

		_, err = run(ctx)
		assert.Error(t, err)
	})

	t.Run("without toml", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)
		ctx.EXPECT().Workdir().Return(root).MaxTimes(2)
		ctx.EXPECT().L().
			DoAndReturn(func() *zap.Logger { return l }).
			AnyTimes()
		ctx.EXPECT().CommitSHA().Return(sha).AnyTimes()

		issues, err := run(ctx)
		require.NoError(t, err)

		for _, issue := range issues {
			info, err := os.Stat(issue.FilePath)
			require.NoError(t, err)
			assert.False(t, info.IsDir())
		}
	})

	t.Run("with root toml", func(t *testing.T) {
		tomlPath := filepath.Join(root, "revive", "fixtures", tomlFileName)

		data, err := os.ReadFile(tomlPath)
		require.NoError(t, err)

		parentTomlPath := filepath.Join(root, "revive.toml")
		f, err := os.OpenFile(parentTomlPath, os.O_RDWR|os.O_CREATE, 0666)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write(data)
		require.NoError(t, err)

		defer os.Remove(parentTomlPath)

		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)
		ctx.EXPECT().Workdir().Return(root).MaxTimes(2)
		ctx.EXPECT().L().
			DoAndReturn(func() *zap.Logger { return l }).
			AnyTimes()
		ctx.EXPECT().CommitSHA().Return(sha).AnyTimes()

		issues, err := run(ctx)
		assert.NoError(t, err)

		for _, issue := range issues {
			info, err := os.Stat(issue.FilePath)
			assert.NoError(t, err)
			assert.False(t, info.IsDir())
		}
	})

	t.Run("with local toml", func(t *testing.T) {
		tomlPath := filepath.Join(root, "revive", "fixtures", tomlFileName)

		data, err := os.ReadFile(tomlPath)
		require.NoError(t, err)

		localTomlPath := filepath.Join(root, "revive", "revive.toml")
		f, err := os.OpenFile(localTomlPath, os.O_RDWR|os.O_CREATE, 0666)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write(data)
		require.NoError(t, err)

		defer os.Remove(localTomlPath)

		ctrl := gomock.NewController(t)
		ctx := sdkmocks.NewMockContext(ctrl)
		ctx.EXPECT().Workdir().Return(root).MaxTimes(2)
		ctx.EXPECT().L().
			DoAndReturn(func() *zap.Logger { return l }).
			AnyTimes()
		ctx.EXPECT().CommitSHA().Return(sha).AnyTimes()

		issues, err := run(ctx)
		assert.NoError(t, err)

		for _, issue := range issues {
			info, err := os.Stat(issue.FilePath)
			assert.NoError(t, err)
			assert.False(t, info.IsDir())
		}
	})
}

func TestReviveTomlExists(t *testing.T) {
	root := common.FindParentDir(t)
	fixtureTomlPath := filepath.Join(root, "revive/fixtures")
	_, ok, err := reviveTomlExists(zap.NewNop(), fixtureTomlPath)
	require.NoError(t, err)
	assert.True(t, ok)
}
