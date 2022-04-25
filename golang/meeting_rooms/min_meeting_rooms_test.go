package meeting_rooms

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	intervals := [][]int{
		{2,36},
		{3,4},
		{13,34},
		{16,20},
		{39,46},
	}
	require.Equal(t, 3, MinMeetingRooms(intervals))
	require.Equal(t, 3, ChronologicalOrderingMeeting(intervals))
}

func TestCase2(t *testing.T) {
	intervals := [][]int{
		{1,8},
		{6,20},
		{9,16},
		{13,17},
	}
	require.Equal(t, 3, MinMeetingRooms(intervals))
	require.Equal(t, 3, ChronologicalOrderingMeeting(intervals))
}

func TestCase3(t *testing.T) {
	intervals := [][]int{
		{1293,2986},
		{848,3846},
		{4284,5907},
		{4466,4781},
		{518,2918},
		{300,5870},
	}
	require.Equal(t, 4, MinMeetingRooms(intervals))
	require.Equal(t, 4, ChronologicalOrderingMeeting(intervals))
}

func TestCase4(t *testing.T) {
	intervals := [][]int{
		{26,29},
		{19,26},
		{19,28},
		{4,19},
		{4,25},
	}
	require.Equal(t, 3, MinMeetingRooms(intervals))
	require.Equal(t, 3, ChronologicalOrderingMeeting(intervals))
}

func TestCase5(t *testing.T) {
	intervals := [][]int{
		{12,13},
		{8,14},
		{6,9},
	}
	require.Equal(t, 2, MinMeetingRooms(intervals))
	require.Equal(t, 2, ChronologicalOrderingMeeting(intervals))
}

func TestCase6(t *testing.T) {
	intervals := [][]int{
		{6,10},
		{13,14},
		{12,14},
	}
	require.Equal(t, 2, MinMeetingRooms(intervals))
	require.Equal(t, 2, ChronologicalOrderingMeeting(intervals))
}

func TestCase7(t *testing.T) {
	intervals := [][]int{
		{1,5},
		{8,9},
		{8,9},
	}
	require.Equal(t, 2, MinMeetingRooms(intervals))
	require.Equal(t, 2, ChronologicalOrderingMeeting(intervals))
}

func TestCase8(t *testing.T) {
	intervals := [][]int{
		{2,15},
		{36,45},
		{9,29},
		{16,23},
		{4,9},
	}
	require.Equal(t, 2, MinMeetingRooms(intervals))
	require.Equal(t, 2, ChronologicalOrderingMeeting(intervals))
}
