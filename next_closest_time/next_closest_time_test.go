package next_closest_time

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	PrintAllNextClosestTime()
	os.Exit(exitVal)
}

func TestNextClosestTime(t *testing.T) {
	time := "19:34"
	require.Equal(t, "19:39", NextClosestTime(time))

	time = "23:55"
	require.Equal(t, "22:22", NextClosestTime(time))

	time = "23:54"
	require.Equal(t, "23:55", NextClosestTime(time))

	time = "01:02"
	require.Equal(t, "01:10", NextClosestTime(time))

	time = "01:03"
	require.Equal(t, "01:10", NextClosestTime(time))

	time = "13:55"
	require.Equal(t, "15:11", NextClosestTime(time))

	time = "20:48"
	require.Equal(t, "22:00", NextClosestTime(time))

	time = "15:55"
	require.Equal(t, "11:11", NextClosestTime(time))
}

func PrintAllNextClosestTime() {
	for i := 0; i < 24; i++ {
		for j := 0; j < 60; j++ {
			hour := fmt.Sprintf("%02d", i)
			minute := fmt.Sprintf("%02d", j)
			input := hour + ":" + minute

			fmt.Println(input, "->", NextClosestTime(input))
		}
	}
}
