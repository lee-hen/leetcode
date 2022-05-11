package strobogrammatic_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []string{"100001", "600009", "800008", "900006", "110011", "610019", "810018", "910016", "160091", "660099", "860098", "960096", "180081", "680089", "880088", "980086", "190061", "690069", "890068", "990066", "101101", "601109", "801108", "901106", "111111", "611119", "811118", "911116", "161191", "661199", "861198", "961196", "181181", "681189", "881188", "981186", "191161", "691169", "891168", "991166", "106901", "606909", "806908", "906906", "116911", "616919", "816918", "916916", "166991", "666999", "866998", "966996", "186981", "686989", "886988", "986986", "196961", "696969", "896968", "996966", "108801", "608809", "808808", "908806", "118811", "618819", "818818", "918816", "168891", "668899", "868898", "968896", "188881", "688889", "888888", "988886", "198861", "698869", "898868", "998866", "109601", "609609", "809608", "909606", "119611", "619619", "819618", "919616", "169691", "669699", "869698", "969696", "189681", "689689", "889688", "989686", "199661", "699669", "899668", "999666"}
	output := FindStrobogrammatic(6)
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []string{"10001", "60009", "80008", "90006", "11011", "61019", "81018", "91016", "16091", "66099", "86098", "96096", "18081", "68089", "88088", "98086", "19061", "69069", "89068", "99066", "10101", "60109", "80108", "90106", "11111", "61119", "81118", "91116", "16191", "66199", "86198", "96196", "18181", "68189", "88188", "98186", "19161", "69169", "89168", "99166", "10801", "60809", "80808", "90806", "11811", "61819", "81818", "91816", "16891", "66899", "86898", "96896", "18881", "68889", "88888", "98886", "19861", "69869", "89868", "99866"}
	output := FindStrobogrammatic(5)
	require.Equal(t, expected, output)
}

func TestCase1_1(t *testing.T) {
	expected := []string{"100001","101101","106901","108801","109601","110011","111111","116911","118811","119611","160091","161191","166991","168891","169691","180081","181181","186981","188881","189681","190061","191161","196961","198861","199661","600009","601109","606909","608809","609609","610019","611119","616919","618819","619619","660099","661199","666999","668899","669699","680089","681189","686989","688889","689689","690069","691169","696969","698869","699669","800008","801108","806908","808808","809608","810018","811118","816918","818818","819618","860098","861198","866998","868898","869698","880088","881188","886988","888888","889688","890068","891168","896968","898868","899668","900006","901106","906906","908806","909606","910016","911116","916916","918816","919616","960096","961196","966996","968896","969696","980086","981186","986986","988886","989686","990066","991166","996966","998866","999666"}
	output := findStrobogrammatic(6)
	require.Equal(t, expected, output)
}

func TestCase2_1(t *testing.T) {
	expected := []string{"10001","10101","10801","11011","11111","11811","16091","16191","16891","18081","18181","18881","19061","19161","19861","60009","60109","60809","61019","61119","61819","66099","66199","66899","68089","68189","68889","69069","69169","69869","80008","80108","80808","81018","81118","81818","86098","86198","86898","88088","88188","88888","89068","89168","89868","90006","90106","90806","91016","91116","91816","96096","96196","96896","98086","98186","98886","99066","99166","99866"}

	output := findStrobogrammatic(5)
	require.Equal(t, expected, output)
}