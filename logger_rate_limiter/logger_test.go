package logger_rate_limiter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogger_ShouldPrintMessage(t *testing.T) {
	obj := Constructor()

	require.True(t, obj.ShouldPrintMessage(1, "foo"))
	require.True(t, obj.ShouldPrintMessage(2, "bar"))
	require.False(t, obj.ShouldPrintMessage(3, "foo"))
	require.False(t, obj.ShouldPrintMessage(8, "bar"))
	require.False(t, obj.ShouldPrintMessage(10, "foo"))
	require.True(t, obj.ShouldPrintMessage(11, "foo"))
}

