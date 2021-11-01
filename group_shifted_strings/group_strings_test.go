package group_shifted_strings

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestGroupStrings(t *testing.T) {
	inputs := [][]string{
		{"lno","izwjzqpponmxwdf","yqxxgjvznxywrswwrjmkn","rifsizyyxwvgfmo","inacgfrmrljgudeetdtvhrcy","asgjbocudyw","jwalbmsoggfwmqnwc","dxabowpgmjdepw","qzwjyjguw","rdkfgekpcbsowhhpomrd","elhgngfmhxkpykxckwmgnrst","skuua","kqjoykcgxatwgpkezowbopf","dpistruwltbdb","rcevdozqhpbjgrrkaxkjdwno","hgydnlweo","klrsskgvlokxzwcikrsgzql","uuoimefyaxvpivbkjjxyo","fbdfqiu","qsstbmbchhmevvyqyz","xghx","bngqrpsujrzbz","iadvsiqnxyjyukujncpqy","ah","zleertfesuwmg","rfpjylhaqdgnyikwlgakpthz","kjbgqozhr","rbwqdyhbtfjfvwpozkbdur","xtvxiam","ydgylser","nwtgvgdrt","fknfszly","nqlvoohpf","eectpfhirrjpaqbmbjhl","nnhbfxyrtqoiboudccqrh","uismbokdtgjqblnzojdnswkc","ldrumznfojh","hnglvhzduxqtdmhbwltylmc","zwrmbabfoaono","hoiepqudvgqvvjasboew","phkczpxuefqfbrbqujwxf","ltgqckhxdabumn","jqdxigjjsrafr","kkbdemrkgywfxyupgxgjlbv","yhvlpdiawozdvaxs","hmthkrfdze","arjxxsrdfmspqqqydauewth","vytdwwpxn","ouhnffqhdbxpzbyzth","zhueqyvlropiab","rxvr","phrrx","zjeylgpjbnrndexwhsjlcz","vcpjusvvedmrd","grtksdofweqyvggzpmzyslcd","ehvqxypwhcackc","eevxygleasqzrsojaradfvp","eqye","utwosauvzuaktlxwunqh","ggnpfdunlpioiowlsd","clzpthmeasdhzebw","jq","qpskowqrvqwgphtsqjmd","zconhjzdgpvszqepuzpre","qcjefdjobarnvggonlqc","pxqyxgbvjlgly","zaghhzvkadzmolrxzghvofa","cirssnq","cvgnuewkmgsjp","cpteuflhzzypfjgpv","coiwbvrx","ejwycbninhfcqzaapzprdnyu","gohpoxsmacxcp","qcwkpjfl","vsnixwxbkwkjk","bbikaypigkdjdjrgny","augsnqtsouzwfzapssm","wdxtefjskvfkkyphqdtl","icfgtbulroijub","wcpvnnypljfxhjghbp","biedkdcjeuhmvhuzhtjdkopq","ofxblentpbige","isiyr","enoe","fhhiqbqrwwbtkknfno","ekie","zqimwpyeamtrp","ewddmpbftdecxyccxpsqt","xza","vpbnilonjpurauvknnh","bhqrrmp","rrpgcsuveewcndozowuy","hkytabszkfdfnf","rkvcjtlzbvhye","tyftwdrplq","jvdj","eqjjwykjxzbrl","hkwvprhloxdahymxchxzm","wgwmf","pgymmhgsubhefffnspjtliw"},
		{"fpbnsbrkbcyzdmmmoisaa","cpjtwqcdwbldwwrryuclcngw","a","fnuqwejouqzrif","js","qcpr","zghmdiaqmfelr","iedda","l","dgwlvcyubde","lpt","qzq","zkddvitlk","xbogegswmad","mkndeyrh","llofdjckor","lebzshcb","firomjjlidqpsdeqyn","dclpiqbypjpfafukqmjnjg","lbpabjpcmkyivbtgdwhzlxa","wmalmuanxvjtgmerohskwil","yxgkdlwtkekavapflheieb","oraxvssurmzybmnzhw","ohecvkfe","kknecibjnq","wuxnoibr","gkxpnpbfvjm","lwpphufxw","sbs","txb","ilbqahdzgij","i","zvuur","yfglchzpledkq","eqdf","nw","aiplrzejplumda","d","huoybvhibgqibbwwdzhqhslb","rbnzendwnoklpyyyauemm"},
	}
	expects := [][][]string {
		[][]string{
			{"izwjzqpponmxwdf", "rifsizyyxwvgfmo"},
			{"vsnixwxbkwkjk", "zwrmbabfoaono"},
			{"bngqrpsujrzbz", "dpistruwltbdb"},
			{"clzpthmeasdhzebw", "yhvlpdiawozdvaxs"},
			{"eectpfhirrjpaqbmbjhl", "rrpgcsuveewcndozowuy"},
			{"hmthkrfdze", "tyftwdrplq"},
			{"lno", "xza"},
			{"hoiepqudvgqvvjasboew", "wdxtefjskvfkkyphqdtl"},
			{"ekie", "rxvr"},
			{"bhqrrmp", "cirssnq"},
			{"qpskowqrvqwgphtsqjmd", "utwosauvzuaktlxwunqh"},
			{"gohpoxsmacxcp", "pxqyxgbvjlgly"},
			{"cvgnuewkmgsjp", "rkvcjtlzbvhye"},
			{"hgydnlweo", "kjbgqozhr"},
			{"klrsskgvlokxzwcikrsgzql", "zaghhzvkadzmolrxzghvofa"},
			{"augsnqtsouzwfzapssm", "vpbnilonjpurauvknnh"},
			{"eevxygleasqzrsojaradfvp", "kkbdemrkgywfxyupgxgjlbv"},
			{"enoe", "xghx"},
			{"ah", "jq"},
			{"qcjefdjobarnvggonlqc", "rdkfgekpcbsowhhpomrd"},
			{"ouhnffqhdbxpzbyzth", "wcpvnnypljfxhjghbp"},
			{"jqdxigjjsrafr", "vcpjusvvedmrd"},
			{"nwtgvgdrt", "qzwjyjguw"},
			{"hkwvprhloxdahymxchxzm", "zconhjzdgpvszqepuzpre"},
			{"dxabowpgmjdepw", "icfgtbulroijub"},
			{"phrrx", "skuua"},
			{"hnglvhzduxqtdmhbwltylmc", "kqjoykcgxatwgpkezowbopf"},
			{"ofxblentpbige", "zqimwpyeamtrp"},
			{"nnhbfxyrtqoiboudccqrh", "uuoimefyaxvpivbkjjxyo"},
			{"rbwqdyhbtfjfvwpozkbdur", "zjeylgpjbnrndexwhsjlcz"},
			{"cpteuflhzzypfjgpv", "jwalbmsoggfwmqnwc"},
			{"fbdfqiu", "xtvxiam"},
			{"grtksdofweqyvggzpmzyslcd", "rcevdozqhpbjgrrkaxkjdwno"},
			{"fhhiqbqrwwbtkknfno", "qsstbmbchhmevvyqyz"},
			{"nqlvoohpf", "vytdwwpxn"},
			{"bbikaypigkdjdjrgny", "ggnpfdunlpioiowlsd"},
			{"iadvsiqnxyjyukujncpqy", "phkczpxuefqfbrbqujwxf"},
			{"biedkdcjeuhmvhuzhtjdkopq", "elhgngfmhxkpykxckwmgnrst"},
			{"rfpjylhaqdgnyikwlgakpthz", "uismbokdtgjqblnzojdnswkc"},
			{"coiwbvrx", "qcwkpjfl"},
			{"ejwycbninhfcqzaapzprdnyu", "inacgfrmrljgudeetdtvhrcy"},
			{"isiyr", "wgwmf"},
			{"ltgqckhxdabumn", "zhueqyvlropiab"},
			{"asgjbocudyw", "ldrumznfojh"},
			{"arjxxsrdfmspqqqydauewth", "pgymmhgsubhefffnspjtliw"},
			{"eqjjwykjxzbrl", "zleertfesuwmg"},
			{"ehvqxypwhcackc", "hkytabszkfdfnf"},
			{"fknfszly", "ydgylser"},
			{"ewddmpbftdecxyccxpsqt", "yqxxgjvznxywrswwrjmkn"},
			{"eqye", "jvdj"},
		},
		[][]string{
			{"a","d","i","l"},
			{"eqdf","qcpr"},
			{"lpt","txb"},
			{"yfglchzpledkq","zghmdiaqmfelr"},
			{"kknecibjnq","llofdjckor"},
			{"cpjtwqcdwbldwwrryuclcngw","huoybvhibgqibbwwdzhqhslb"},
			{"lbpabjpcmkyivbtgdwhzlxa","wmalmuanxvjtgmerohskwil"},
			{"iedda","zvuur"},
			{"js","nw"},
			{"lebzshcb","ohecvkfe"},
			{"dgwlvcyubde","ilbqahdzgij"},
			{"lwpphufxw","zkddvitlk"},
			{"qzq","sbs"},
			{"dclpiqbypjpfafukqmjnjg","yxgkdlwtkekavapflheieb"},
			{"mkndeyrh","wuxnoibr"},
			{"firomjjlidqpsdeqyn","oraxvssurmzybmnzhw"},
			{"gkxpnpbfvjm","xbogegswmad"},
			{"fpbnsbrkbcyzdmmmoisaa","rbnzendwnoklpyyyauemm"},
			{"aiplrzejplumda","fnuqwejouqzrif"},
		},
	}


	for i := 0; i < 2; i++ {
		actual := GroupStrings(inputs[i])
		expected := expects[i]
		SortStrings(expected)
		SortStrings(actual)
		require.ElementsMatch(t, expected, actual)
		actual = GroupStringsBruteForce(inputs[i])
		SortStrings(actual)
		require.ElementsMatch(t, expected, actual)
	}
}

func SortStrings(strs [][]string) {
	for _, letters := range strs {
		sort.Strings(letters)
	}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
}
