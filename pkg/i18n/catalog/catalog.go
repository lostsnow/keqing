// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_US":   &dictionary{index: en_USIndex, data: en_USData},
		"zh_Hans": &dictionary{index: zh_HansIndex, data: zh_HansData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"Albedo":                  68,
	"Alhaitham":               50,
	"Aloy":                    60,
	"Amber":                   7,
	"Anemo":                   74,
	"Arataki Itto":            69,
	"Barbara":                 17,
	"Beidou":                  37,
	"Bennett":                 4,
	"Candace":                 22,
	"Chongyun":                52,
	"Collei":                  48,
	"Cryo":                    77,
	"Cut to the chase":        0,
	"Cyno":                    45,
	"Dehya":                   14,
	"Dendro":                  76,
	"Diluc":                   8,
	"Diona":                   55,
	"Dori":                    44,
	"Electro":                 75,
	"Eula":                    58,
	"Faruzan":                 32,
	"Fischl":                  36,
	"Ganyu":                   56,
	"Geo":                     78,
	"Gorou":                   70,
	"Hu Tao":                  10,
	"Hydro":                   73,
	"Jean":                    27,
	"Kaedehara Kazuha":        29,
	"Kaeya":                   54,
	"Kamisato Ayaka":          59,
	"Kamisato Ayato":          20,
	"Keqing":                  35,
	"Klee":                    5,
	"Kujou Sara":              40,
	"Kuki Shinobu":            43,
	"Layla":                   62,
	"Lisa":                    39,
	"Mika":                    63,
	"Mona":                    15,
	"Nahida":                  49,
	"Nilou":                   23,
	"Ningguang":               66,
	"No such character":       1,
	"No such character guide": 2,
	"Noelle":                  65,
	"Pyro":                    72,
	"Qiqi":                    53,
	"Raiden Shogun":           41,
	"Razor":                   38,
	"Rosaria":                 57,
	"Sangonomiya Kokomi":      19,
	"Sayu":                    30,
	"Shenhe":                  61,
	"Shikanoin Heizou":        31,
	"Sucrose":                 25,
	"Tartaglia":               18,
	"Thoma":                   13,
	"Tighnari":                47,
	"Traveler Anemo":          24,
	"Traveler Dendro":         46,
	"Traveler Electro":        34,
	"Traveler Geo":            64,
	"Venti":                   26,
	"Wanderer":                33,
	"What are you looking for, it does not exist": 3,
	"Xiangling": 6,
	"Xiao":      28,
	"Xingqiu":   16,
	"Xinyan":    9,
	"Yae Miko":  42,
	"Yanfei":    11,
	"Yaoyao":    51,
	"Yelan":     21,
	"Yoimiya":   12,
	"Yun Jin":   71,
	"Zhongli":   67,
}

var en_USIndex = []uint32{ // 80 elements
	// Entry 0 - 1F
	0x00000000, 0x00000011, 0x00000023, 0x0000003b,
	0x00000067, 0x0000006f, 0x00000074, 0x0000007e,
	0x00000084, 0x0000008a, 0x00000091, 0x00000098,
	0x0000009f, 0x000000a7, 0x000000ad, 0x000000b3,
	0x000000b8, 0x000000c0, 0x000000c8, 0x000000d2,
	0x000000e5, 0x000000f4, 0x000000fa, 0x00000102,
	0x00000108, 0x00000117, 0x0000011f, 0x00000125,
	0x0000012a, 0x0000012f, 0x00000140, 0x00000145,
	// Entry 20 - 3F
	0x00000156, 0x0000015e, 0x00000167, 0x00000178,
	0x0000017f, 0x00000186, 0x0000018d, 0x00000193,
	0x00000198, 0x000001a3, 0x000001b1, 0x000001ba,
	0x000001c7, 0x000001cc, 0x000001d1, 0x000001e1,
	0x000001ea, 0x000001f1, 0x000001f8, 0x00000202,
	0x00000209, 0x00000212, 0x00000217, 0x0000021d,
	0x00000223, 0x00000229, 0x00000231, 0x00000236,
	0x00000245, 0x0000024a, 0x00000251, 0x00000257,
	// Entry 40 - 5F
	0x0000025c, 0x00000269, 0x00000270, 0x0000027a,
	0x00000282, 0x00000289, 0x00000296, 0x0000029c,
	0x000002a4, 0x000002a9, 0x000002af, 0x000002b5,
	0x000002bd, 0x000002c4, 0x000002c9, 0x000002cd,
} // Size: 344 bytes

const en_USData string = "" + // Size: 717 bytes
	"\x02Cut to the chase\x02No such character\x02No such character guide\x02" +
	"What are you looking for, it does not exist\x02Bennett\x02Klee\x02Xiangl" +
	"ing\x02Amber\x02Diluc\x02Xinyan\x02Hu Tao\x02Yanfei\x02Yoimiya\x02Thoma" +
	"\x02Dehya\x02Mona\x02Xingqiu\x02Barbara\x02Tartaglia\x02Sangonomiya Koko" +
	"mi\x02Kamisato Ayato\x02Yelan\x02Candace\x02Nilou\x02Traveler Anemo\x02S" +
	"ucrose\x02Venti\x02Jean\x02Xiao\x02Kaedehara Kazuha\x02Sayu\x02Shikanoin" +
	" Heizou\x02Faruzan\x02Wanderer\x02Traveler Electro\x02Keqing\x02Fischl" +
	"\x02Beidou\x02Razor\x02Lisa\x02Kujou Sara\x02Raiden Shogun\x02Yae Miko" +
	"\x02Kuki Shinobu\x02Dori\x02Cyno\x02Traveler Dendro\x02Tighnari\x02Colle" +
	"i\x02Nahida\x02Alhaitham\x02Yaoyao\x02Chongyun\x02Qiqi\x02Kaeya\x02Diona" +
	"\x02Ganyu\x02Rosaria\x02Eula\x02Kamisato Ayaka\x02Aloy\x02Shenhe\x02Layl" +
	"a\x02Mika\x02Traveler Geo\x02Noelle\x02Ningguang\x02Zhongli\x02Albedo" +
	"\x02Arataki Itto\x02Gorou\x02Yun Jin\x02Pyro\x02Hydro\x02Anemo\x02Electr" +
	"o\x02Dendro\x02Cryo\x02Geo"

var zh_HansIndex = []uint32{ // 80 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000d, 0x0000001a, 0x00000038,
	0x00000053, 0x0000005d, 0x00000064, 0x0000006b,
	0x00000072, 0x0000007c, 0x00000083, 0x0000008a,
	0x00000091, 0x00000098, 0x0000009f, 0x000000a9,
	0x000000b0, 0x000000b7, 0x000000c1, 0x000000ce,
	0x000000de, 0x000000eb, 0x000000f2, 0x000000fc,
	0x00000103, 0x00000112, 0x00000119, 0x00000120,
	0x00000124, 0x00000128, 0x00000135, 0x0000013c,
	// Entry 20 - 3F
	0x0000014c, 0x00000156, 0x00000160, 0x0000016f,
	0x00000176, 0x00000180, 0x00000187, 0x0000018e,
	0x00000195, 0x000001a2, 0x000001af, 0x000001bc,
	0x000001c6, 0x000001cd, 0x000001d4, 0x000001e3,
	0x000001ed, 0x000001f4, 0x000001fe, 0x0000020b,
	0x00000212, 0x00000219, 0x00000220, 0x00000227,
	0x00000231, 0x00000238, 0x00000245, 0x0000024c,
	0x00000259, 0x00000263, 0x0000026a, 0x00000274,
	// Entry 40 - 5F
	0x0000027b, 0x0000028a, 0x00000294, 0x0000029b,
	0x000002a2, 0x000002ac, 0x000002b9, 0x000002c0,
	0x000002c7, 0x000002cb, 0x000002cf, 0x000002d3,
	0x000002d7, 0x000002db, 0x000002df, 0x000002e3,
} // Size: 344 bytes

const zh_HansData string = "" + // Size: 739 bytes
	"\x02斩尽牛杂\x02查无此人\x02此人很神秘, 没有情报\x02你找啥呢, 不存在的\x02班尼特\x02可莉\x02香菱\x02安柏" +
	"\x02迪卢克\x02辛焱\x02胡桃\x02烟绯\x02宵宫\x02托马\x02迪希雅\x02莫娜\x02行秋\x02芭芭拉\x02达达利亚" +
	"\x02珊瑚宫心海\x02神里绫人\x02夜兰\x02坎蒂丝\x02妮露\x02旅行者(风)\x02砂糖\x02温迪\x02琴\x02魈\x02" +
	"枫原万叶\x02早柚\x02鹿野院平藏\x02珐露珊\x02流浪者\x02旅行者(雷)\x02刻晴\x02菲谢尔\x02北斗\x02雷泽" +
	"\x02丽莎\x02九条裟罗\x02雷电将军\x02八重神子\x02久岐忍\x02多莉\x02赛诺\x02旅行者(草)\x02提纳里\x02柯莱" +
	"\x02纳西妲\x02艾尔海森\x02瑶瑶\x02重云\x02七七\x02凯亚\x02迪奥娜\x02甘雨\x02罗莎莉亚\x02优菈\x02神里" +
	"绫华\x02埃洛伊\x02申鹤\x02莱依拉\x02米卡\x02旅行者(岩)\x02诺艾尔\x02凝光\x02钟离\x02阿贝多\x02荒泷" +
	"一斗\x02五郎\x02云堇\x02火\x02水\x02风\x02雷\x02草\x02冰\x02岩"

	// Total table size 2144 bytes (2KiB); checksum: 778099D1
