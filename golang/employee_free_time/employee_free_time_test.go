package employee_free_time

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	schedule := [][]*Interval{
		{
			&Interval{
				1, 2,
			},
			&Interval{
				5, 6,
			},
		},
		{
			&Interval{
				1, 3,
			},
			&Interval{
				4, 10,
			},
		},
	}

	expect := []*Interval{
		{
			3, 4,
		},
	}

	commonFreeIntervals := EmployeeFreeTime(schedule)
	commFreeIntervals := employeeFreeTime(schedule)

	for i, interval := range commonFreeIntervals {
		require.Equal(t, expect[i].Start, interval.Start)
		require.Equal(t, expect[i].End, interval.End)

		require.Equal(t, expect[i].Start, commFreeIntervals[i].Start)
		require.Equal(t, expect[i].End, commFreeIntervals[i].End)
	}
}

func TestCase2(t *testing.T) {
	schedule := [][]*Interval{
		{
			&Interval{
				1, 3,
			},
			&Interval{
				6, 7,
			},
		},
		{
			&Interval{
				2, 4,
			},
		},
		{
			&Interval{
				2, 5,
			},
			&Interval{
				9, 12,
			},
		},
	}

	expect := []*Interval{
		{
			5, 6,
		},
		{
			7, 9,
		},
	}

	commonFreeIntervals := EmployeeFreeTime(schedule)
	commFreeIntervals := employeeFreeTime(schedule)

	for i, interval := range commonFreeIntervals {
		require.Equal(t, expect[i].Start, interval.Start)
		require.Equal(t, expect[i].End, interval.End)

		require.Equal(t, expect[i].Start, commFreeIntervals[i].Start)
		require.Equal(t, expect[i].End, commFreeIntervals[i].End)
	}
}

func TestCase3(t *testing.T) {
	schedule := [][]*Interval{
		{
			&Interval{
				1, 3,
			},
			&Interval{
				6, 7,
			},
		},
		{
			&Interval{
				2, 11,
			},
		},
		{
			&Interval{
				2, 5,
			},
			&Interval{
				9, 12,
			},
		},
	}
	commonFreeIntervals := EmployeeFreeTime(schedule)
	require.Equal(t, []*Interval{}, commonFreeIntervals)

	commFreeIntervals := employeeFreeTime(schedule)
	require.Equal(t, []*Interval{}, commFreeIntervals)
}

func TestCase4(t *testing.T) {
	schedule := [][]*Interval{
		{
			&Interval{
				1, 2,
			},
			&Interval{
				5, 6,
			},
		},
		{
			&Interval{
				1, 3,
			},
		},
		{
			&Interval{
				4, 10,
			},
		},
	}

	expect := []*Interval{
		{
			3, 4,
		},
	}

	commonFreeIntervals := EmployeeFreeTime(schedule)
	commFreeIntervals := employeeFreeTime(schedule)

	for i, interval := range commonFreeIntervals {
		require.Equal(t, expect[i].Start, interval.Start)
		require.Equal(t, expect[i].End, interval.End)

		require.Equal(t, expect[i].Start, commFreeIntervals[i].Start)
		require.Equal(t, expect[i].End, commFreeIntervals[i].End)
	}
}
