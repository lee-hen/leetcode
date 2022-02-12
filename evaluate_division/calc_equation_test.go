package evaluate_division

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}

	expected := []float64{6.0, 0.5, -1.0, 1.0, -1.0}
	actual := CalcEquation(equations, values, queries)
	require.Equal(t, expected, actual)

	equations = [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values = []float64{1.5, 2.5, 5.0}
	queries = [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}

	expected = []float64{3.75000, 0.40000, 5.00000, 0.20000}
	actual = CalcEquation(equations, values, queries)
	require.Equal(t, expected, actual)

	equations = [][]string{{"a", "b"}}
	values = []float64{0.5}
	queries = [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}

	expected = []float64{0.50000, 2.00000, -1.00000, -1.00000}
	actual = CalcEquation(equations, values, queries)
	require.Equal(t, expected, actual)

	equations = [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}
	values = []float64{3.0, 4.0, 5.0, 6.0}
	queries = [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}

	expected = []float64{360.00000, 0.008333333333333333, 20.00000, 1.00000, -1.00000, -1.00000}
	actual = CalcEquation(equations, values, queries)
	require.Equal(t, expected, actual)

	equations = [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x1", "x4"}, {"x2", "x5"}}
	values = []float64{3.0, 0.5, 3.4, 5.6}
	queries = [][]string{{"x2","x4"}, {"x1","x5"}, {"x1","x3"}, {"x5","x5"}, {"x5","x1"}, {"x3","x4"},{"x4","x3"},{"x6","x6"},{"x0","x0"}}

	expected = []float64{1.1333333333333333, 16.799999999999997, 1.5, 1, 0.05952380952380952, 2.2666666666666666, 0.4411764705882353, -1, -1}
	actual = CalcEquation(equations, values, queries)
	require.Equal(t, expected, actual)
}
