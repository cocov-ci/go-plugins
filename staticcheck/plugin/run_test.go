package plugin

import (
	"fmt"
	"os"
	"testing"

	"github.com/cocov-ci/go-plugins/common"
	sdkmocks "github.com/cocov-ci/go-plugins/common/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestRun(t *testing.T) {
	root := common.FindParentDir(t)
	sha := "sha"

	ctrl := gomock.NewController(t)
	ctx := sdkmocks.NewMockContext(ctrl)

	l := zap.NewNop()
	ctx.EXPECT().Workdir().Return(root)
	ctx.EXPECT().L().
		DoAndReturn(func() *zap.Logger { return l }).
		AnyTimes()

	ctx.EXPECT().CommitSHA().Return(sha).AnyTimes()

	issues, err := run(ctx)
	require.NoError(t, err)

	for _, i := range issues {
		fmt.Println(i)
		f, err := os.Stat(i.FilePath)
		require.NoError(t, err)
		require.False(t, f.IsDir())
	}
}
