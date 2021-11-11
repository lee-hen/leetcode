package meeting_rooms

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCanAttendMeetings(t *testing.T) {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	require.False(t, CanAttendMeetings(intervals))

	intervals = [][]int{{7,10}, {2, 4}}
	require.True(t, CanAttendMeetings(intervals))
}
