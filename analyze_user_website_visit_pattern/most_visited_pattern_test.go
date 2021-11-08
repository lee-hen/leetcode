package analyze_user_website_visit_pattern

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	username := []string{"joe","joe","joe","james","james","james","james","mary","mary","mary"}
	timestamp := []int{1,2,3,4,5,6,7,8,9,10}
	website := []string{"home","about","career","home","cart","maps","home","home","about","career"}

	require.Equal(t, []string{"home","about","career"}, MostVisitedPattern(username, timestamp, website))
}

func TestCase2(t *testing.T) {
	username := []string{"ua","ua","ua","ub","ub","ub"}
	timestamp := []int{1,2,3,4,5,6}
	website := []string{"a","b","c","a","b","a"}

	require.Equal(t, []string{"a", "b", "a"}, MostVisitedPattern(username, timestamp, website))
}

func TestCase3(t *testing.T) {
	username := []string{"dowg","dowg","dowg"}
	timestamp := []int{158931262,562600350,148438945}
	website := []string{"y","loedo","y"}

	require.Equal(t, []string{"y","y","loedo"}, MostVisitedPattern(username, timestamp, website))
}

func TestCase4(t *testing.T) {
	username := []string{"h","eiy","cq","h","cq","txldsscx","cq","txldsscx","h","cq","cq"}
	timestamp := []int{527896567,334462937,517687281,134127993,859112386,159548699,51100299,444082139,926837079,317455832,411747930}
	website := []string{"hibympufi","hibympufi","hibympufi","hibympufi","hibympufi","hibympufi","hibympufi","hibympufi","yljmntrclw","hibympufi","yljmntrclw"}

	require.Equal(t, []string{"hibympufi","hibympufi","yljmntrclw"}, MostVisitedPattern(username, timestamp, website))
}

func TestCase5(t *testing.T) {
	username := []string{"pdtkrjhd","pdtkrjhd","pdtkrjhd","pdtkrjhd","pdtkrjhd","pdtkrjhd"}
	timestamp := []int{210984153,262799291,958396687,605779010,373702273,205190519}
	website := []string{"xgriygexlk","qs","rugydl","bkrok","canlv","cahgsobjjs"}

	require.Equal(t, []string{"cahgsobjjs","bkrok","rugydl"}, MostVisitedPattern(username, timestamp, website))
}

func TestCase6(t *testing.T) {
	username := []string{"c","c","c","c","c","c","c","c","c","c","c","c","c"}
	timestamp := []int{192769792,207063377,333805625,890700372,64199933,245678250,69530300,833864121,527969074,492534187,49153037,143138463,163274379}
	website := []string{"jsips","zkamv","osajva","bxbldeqhz","zkamv","osajva","zkamv","osajva","zkamv","zkamv","zkamv","zkamv","jsips"}

	require.Equal(t, []string{"jsips","jsips","bxbldeqhz"}, MostVisitedPattern(username, timestamp, website))
}

func Test_generateTimestampPatterns(t *testing.T) {
	expect := [][]int{{4, 5, 6}, {4, 5, 7}, {4, 5, 8}, {4, 6, 7}, {4, 6, 8}, {4, 7, 8}, {5, 6, 7}, {5, 6, 8}, {5, 7, 8}, {6, 7, 8}}
	require.Equal(t, expect, generateTimestampPatterns([]int{4, 5, 6, 7, 8}))
}


