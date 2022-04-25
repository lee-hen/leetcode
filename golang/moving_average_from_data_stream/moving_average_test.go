package moving_average_from_data_stream

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMovingAverage_Next(t *testing.T) {
	movingAverage := Constructor(3)

	require.Equal(t, 1.0, movingAverage.Next(1))
	require.Equal(t, 5.5, movingAverage.Next(10))
	require.Equal(t, 4.666666666666667, movingAverage.Next(3))
	require.Equal(t, 6.0, movingAverage.Next(5))
}
