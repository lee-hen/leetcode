package find_and_replace_in_string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindReplaceString(t *testing.T) {
	s := "abcd"
	indices := []int{0, 2}
	sources := []string{"a", "cd"}
	targets := []string{"eee", "ffff"}
	require.Equal(t, "eeebffff", FindReplaceString(s, indices, sources, targets))

	s = "abcd"
	indices = []int{0, 2}
	sources = []string{"ab", "ec"}
	targets = []string{"eee", "ffff"}
	require.Equal(t, "eeecd", FindReplaceString(s, indices, sources, targets))

	s = "abc"
	indices = []int{0, 1}
	sources = []string{"ab", "bc"}
	targets = []string{"eee", "ffff"}
	require.Equal(t, "", FindReplaceString(s, indices, sources, targets))

	s = "vmokgggqzp"
	indices = []int{3, 5, 1}
	sources = []string{"kg", "ggq", "mo"}
	targets = []string{"s", "so", "bfr"}
	require.Equal(t, "vbfrssozp", FindReplaceString(s, indices, sources, targets))

	s = "jjievdtjfb"
	indices = []int{4, 6, 1}
	sources = []string{"md", "tjgb", "jf"}
	targets = []string{"foe", "oov", "e"}
	require.Equal(t, "jjievdtjfb", FindReplaceString(s, indices, sources, targets))

	s = "mhnbzxkwzxtaanmhtoirxheyanoplbvjrovzudznmetkkxrdmr"
	indices = []int{46, 29, 2, 44, 31, 26, 42, 9, 38, 23, 36, 12, 16, 7, 33, 18}
	sources = []string{"rym", "kv", "nbzxu", "vx", "js", "tp", "tc", "jta", "zqm", "ya", "uz", "avm", "tz", "wn", "yv", "ird"}
	targets = []string{"gfhc", "uq", "dntkw", "wql", "s", "dmp", "jqi", "fp", "hs", "aqz", "ix", "jag", "n", "l", "y", "zww"}
	require.Equal(t, "mhnbzxkwzxtaanmhtoirxheaqznoplbvjrovzudznmetkkxrdmr", FindReplaceString(s, indices, sources, targets))

	s = "jsrdovvyvdngtrvxkzxkwzirgjsmzaoctjvpncjyxcarcxzymazmnrzatkffhfkalrajemoowzaqfjczgtbcpidunmasgfxdhifw"
	indices = []int{18, 76, 2, 13, 4, 40, 91, 0, 68, 87, 50, 10, 48, 31, 82, 6, 66, 58, 73, 71, 45, 38, 95, 29, 42, 79, 8, 53, 15, 25, 85, 35, 22}
	sources = []string{"xkwd", "oj", "rk", "yv", "rv", "xn", "sgfg", "is", "edo", "uhma", "zr", "nz", "ha", "ztjv", "bu", "vo", "ak", "ffhfta", "ia", "lw", "xz", "jx", "djifw", "zo", "arp", "dg", "cd", "rzzt", "uk", "ls", "id", "pc", "fr"}
	targets = []string{"hen", "h", "qo", "ivk", "hlm", "c", "luey", "dv", "uvd", "whkya", "ea", "vg", "zw", "fmo", "tx", "r", "lof", "wwdzgb", "fxj", "jl", "vi", "id", "edcqlh", "hvy", "zsgf", "rl", "lm", "tmtp", "n", "kvb", "wbl", "d", "ucm"}
	require.Equal(t, "jsrdovvyvdngtrvxkzxkwzirgjsmzaoctjvpncjyxcarcviymazmnrzatkffhfkalrajemoowzaqfjczgtbcpwblunmasgfxdhifw", FindReplaceString(s, indices, sources, targets))

	s = "jcxfgtcezgxblhkdgubhempnaoossyypewihccbzbkdjjxbqvnzqycdlwmrjjfykuitkzfhchuambdagictmjatwnttpcenraowhzmlgfvxcyamfonupldrrnvnebtzqxdjongapktgmiytqiqseizonpitnknfzwuendmvxhbsobidnqwhplolahpijafzjistxtnfdcxxkrruxbphmjqzanebmioyqmyqwayunokwbvmckgpmcxoeqtadafkbxnjvknhmjtlzkiqxirobjcpsikcyhvmoehsompftkxxfkmneqtpjntrcatlwgvgmrrvaraytvhpbidajyqolnzqchxwvpdvchgfnhohypbkzohgdchxspsylhxaefbpzaomwgxghpniyauvvsmcwkkycnxjhfhyrkczzmlyxmyjnxwrijivgndavscpzuwuyibnbzhkitmavsubvdhfblpcjudijvudgqjbscsvyrqnkmhthigprapayxlibvessdqddqmxwtqcipdwouqgrhldfczookpptncyrvvaylktaxjznflzdjorijatuzhdzqmttkmbirsfkvwkrqxscuarxrwcuzvptpucjahdvcepftaidksapahhfjziskrwxafellspalwjvgfrcbswgdtbmibejwtimhdyzyfcgbmdmggmrzrdjbgjerhnuapivvyogxaswrkjuxrxilejjzwqxqthpshriufqqahdrswmpsqcdrmfzharfbduefyspfrsnfsmywygkkidebqnfttsbvfjqoqnxdtwoptampsgxizfgpuiywrtghkguxqcmhksujlcosebfvzewqegxiuxbryqhtynrljfdmsdztatifilkehfarmroxprnntcixzhlireearlfpoddpduplqhjspnnhpsqhiawbnubmpwgkleubldsutnenhugmellxvjrnkulvoyeaihufqlqflwpolekamzjfdnflnlfj"
	indices = []int{644, 593, 508, 216, 367, 211, 585, 319, 39, 345, 665, 968, 570, 755, 722, 900, 578, 406, 855, 599, 326, 135, 451, 282, 993, 43, 660, 951, 118, 304, 240, 918, 739, 380, 138, 158, 354, 689, 822, 544, 78, 306, 713, 827, 484, 418, 513, 671, 475, 743, 844, 815, 232, 291, 610, 22, 442, 199, 556, 786, 54, 799, 980, 401, 439, 928, 962, 595, 615, 29, 603, 13, 489, 502, 887, 63, 391, 302, 19, 170, 415, 700, 558, 629, 533, 681, 335, 102, 870, 468, 850, 881, 93, 781, 454, 895, 132, 258, 273, 85}
	sources = []string{"sp", "sc", "essd", "nebmioyqmyqwayu", "chxsps", "mj", "fkvwkr", "rrv", "zbk", "wvpdvch", "ejwtim", "oyeaihufqlqf", "uzh", "harfbduefy", "jzwqx", "reearl", "tkmbirs", "cn", "ryqht", "rwcu", "tvh", "pk", "zh", "yhvmoe", "nflnl", "jj", "tbmib", "nhugmel", "rrnvnebt", "tp", "gpmcxoeqta", "jspnnh", "ahdr", "bpzaomwgx", "gmiytqiq", "fzwuendm", "nhohypbkzoh", "djbgjerhnua", "ghkg", "cy", "ag", "jntrcatl", "juxrxilej", "xq", "svyrq", "zmlyxmyjnxwrijivgndav", "ddqmxwtqcipd", "hdy", "vudgqjbsc", "swmpsqcdrmfz", "zewqe", "puiy", "okwb", "mpftkxx", "jahdv", "pn", "zuwu", "dcxxkrru", "jz", "fttsbvfjqoqn", "dl", "dtwoptamp", "lwpoleka", "cwkky", "scp", "ia", "rn", "uarx", "cepftaidk", "yy", "zvptp", "hkdg", "nkmhthigpr", "yxlibv", "xprn", "kuitkzfhchu", "pn", "eq", "he", "sobidnqwhplolahpijafzjistx", "kc", "pivvyo", "nflzdjor", "hfjzisk", "dfczoo", "dmgg", "yqolnzqchx", "lgfvxcyamfo", "zta", "pc", "xiu", "farmr", "enraowhz", "debqn", "itmavsu", "xzhl", "nga", "vknhmjtlzk", "objcps", "at"}
	targets = []string{"hre", "eb", "eaos", "txwfmaaawnziac", "weoma", "wis", "gavna", "ydbh", "ki", "mfvmstk", "uixjsxw", "gqmtlrloxukvr", "gcrj", "dycdltutsoy", "xciglm", "vybvqwj", "okadwcwu", "wcn", "qgppaw", "tdv", "obhm", "lsb", "ipm", "ixqvd", "kyhk", "ir", "lmfq", "ibtwuj", "hpszdpv", "d", "gyfpuzrqtaq", "bvxah", "sysiw", "ckrzllzv", "ztyzdpxg", "vmymxildh", "izvjydtdsfg", "bdolsfekorw", "llp", "cp", "no", "ihhzltwwh", "uebjhublp", "ba", "mxzt", "dqcnemigrhisnnqozqdv", "piahpypqhghcy", "tg", "oqdevagd", "cdxvxclnxlfr", "igrt", "wlptw", "dwc", "vryniflm", "tcfp", "c", "nct", "ikalqhj", "po", "xjuhppqtcesci", "v", "jifsgmxb", "ippvxzcbj", "zgpjn", "ef", "mfd", "n", "xqc", "siymapctu", "cnk", "rwzr", "ismpd", "xrhlptgbf", "kmgufwj", "qcj", "pipngvmjcbwk", "g", "s", "cx", "bbjwmjymrngbcjubgltxgejpi", "rd", "efvlpki", "efxyxdmqa", "znvwul", "arqltiq", "ozyoj", "ypgonrmyzvs", "oexeudgxgkbw", "pt", "ujt", "afrt", "robw", "ncttsovm", "qrugjn", "vrpurgiq", "vuaa", "hu", "hjdbdmeoyz", "rsrefef", "nfw"}
	require.Equal(t, "jcxfgtcezgxblismpdubcxmcaoosscnkpewihccbkidirxbqvnzqycvwmrjjfypipngvmjcbwkambdnoictmjnfwwnttpcncttsovmmoexeudgxgkbwnupldhpszdpvzqxdjohulsbtztyzdpxgseizonpitnknvmymxildhvxhbbbjwmjymrngbcjubgltxgejpitnfikalqhjxbphwisqzatxwfmaaawnziacndwcvmckgyfpuzrqtaqdafkbxnjhjdbdmeoyziqxirrsrefefikcixqvdhsovryniflmfkmnsdihhzltwwhwgvgmydbharayobhmpbidajypgonrmyzvsmfvmstkgfizvjydtdsfggdweomaylhxaefckrzllzvghgiyauvvsmzgpjnwcnxjhfhyrrdzdqcnemigrhisnnqozqdvefnctyibnbipmkvrpurgiqbvdhfblujtjudijoqdevagdmxztxrhlptgbfapakmgufwjeaosqpiahpypqhghcywouqgrhlarqltiqkpptncprvvaylktaxpoefxyxdmqaijatgcrjdzqmtokadwcwugavnaqxebxqctdvrwzructcfpsiymapctusapahznvwulrwxafellhrealwjvgfrcbswgdlmfquixjsxwtgzyfcgbmozyojmrzrbdolsfekorwefvlpkigxaswrkuebjhublpxciglmqthpshriufqqsysiwcdxvxclnxlfrdycdltutsoyspfrsnfsmywygkkiqrugjnxjuhppqtcescixjifsgmxbsgxizfgwlptwwrtllpubacmhksujlcosebfvigrtgafrtxbqgppawynrljfdmsdpttifilkehrobwoqcjntcivuaaivybvqwjfpoddpduplqhbvxahpsqhmfdwbnubmpwgkleubldsutneibtwujlxvjnkulvgqmtlrloxukvrippvxzcbjmzjfdkyhkfj", FindReplaceString(s, indices, sources, targets))
}
