package character

import (
	"strings"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/elemental"
	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
	"github.com/lostsnow/keqing/pkg/object"
	"github.com/lostsnow/keqing/pkg/star"
)

type Character struct {
	Id        string
	Elemental *elemental.Elemental
	Names     []string
	Star      int
}

var (
	Bennett   = New("Bennett", elemental.Pyro, star.FOUR, "点赞哥")
	Klee      = New("Klee", elemental.Pyro, star.FIVE, "火花骑士")
	Xiangling = New("Xiangling", elemental.Pyro, star.FOUR)
	Amber     = New("Amber", elemental.Pyro, star.FOUR)
	Diluc     = New("Diluc", elemental.Pyro, star.FIVE, "卢老爷", "正义人")
	Xinyan    = New("Xinyan", elemental.Pyro, star.FOUR)
	HuTao     = New("Hu Tao", elemental.Pyro, star.FIVE, "胡堂主")
	Yanfei    = New("Yanfei", elemental.Pyro, star.FOUR)
	Yoimiya   = New("Yoimiya", elemental.Pyro, star.FIVE)
	Thoma     = New("Thoma", elemental.Pyro, star.FOUR)
	Dehya     = New("Dehya", elemental.Pyro, star.FIVE)

	Mona              = New("Mona", elemental.Hydro, star.FIVE, "梅姬斯图斯卿")
	Xingqiu           = New("Xingqiu", elemental.Hydro, star.FOUR)
	Barbara           = New("Barbara", elemental.Hydro, star.FOUR, "内鬼")
	Tartaglia         = New("Tartaglia", elemental.Hydro, star.FIVE, "公子", "达达鸭", "汤达人", "永别冬都", "枕玉")
	SangonomiyaKokomi = New("Sangonomiya Kokomi", elemental.Hydro, star.FIVE, "观赏鱼")
	KamisatoAyato     = New("Kamisato Ayato", elemental.Hydro, star.FIVE)
	Yelan             = New("Yelan", elemental.Hydro, star.FIVE, "叶天后")
	Candace           = New("Candace", elemental.Hydro, star.FOUR)
	Nilou             = New("Nilou", elemental.Hydro, star.FIVE, "舞娘")

	TravelerAnemo   = New("Traveler Anemo", elemental.Anemo, star.FIVE, "风主")
	Sucrose         = New("Sucrose", elemental.Anemo, star.FOUR)
	Venti           = New("Venti", elemental.Anemo, star.FIVE, "风神", "巴巴托斯", "卖唱的")
	Jean            = New("Jean", elemental.Anemo, star.FIVE, "代理团长", "蒲公英骑士")
	Xiao            = New("Xiao", elemental.Anemo, star.FIVE, "三尺五寸仙人", "降魔大圣", "金鹏大将")
	KaedeharaKazuha = New("Kaedehara Kazuha", elemental.Anemo, star.FIVE, "叶天帝")
	Sayu            = New("Sayu", elemental.Anemo, star.FOUR, "泥头车")
	ShikanoinHeizou = New("Shikanoin Heizou", elemental.Anemo, star.FOUR)
	Faruzan         = New("Faruzan", elemental.Anemo, star.FOUR, "百岁山", "百岁珊")
	Wanderer        = New("Wanderer", elemental.Anemo, star.FIVE, "散兵", "斯卡拉姆齐", "国崩")

	TravelerElectro = New("Traveler Electro", elemental.Electro, star.FIVE, "雷主")
	Keqing          = New("Keqing", elemental.Electro, star.FIVE, "刻师傅", "牛杂师傅", "玉衡星")
	Fischl          = New("Fischl", elemental.Electro, star.FOUR, "断罪之皇女", "小艾咪")
	Beidou          = New("Beidou", elemental.Electro, star.FOUR, "大姐头")
	Razor           = New("Razor", elemental.Electro, star.FOUR, "狼崽子")
	Lisa            = New("Lisa", elemental.Electro, star.FOUR, "图书管理员")
	KujouSara       = New("Kujou Sara", elemental.Electro, star.FOUR, "九条天狗")
	RaidenShogun    = New("Raiden Shogun", elemental.Electro, star.FIVE, "雷神", "巴尔泽布", "雷军", "御建鸣神主尊大御所大人")
	YaeMiko         = New("Yae Miko", elemental.Electro, star.FIVE, "屑狐狸")
	KukiShinobu     = New("Kuki Shinobu", elemental.Electro, star.FOUR, "97忍", "阿忍")
	Dori            = New("Dori", elemental.Electro, star.FOUR)
	Cyno            = New("Cyno", elemental.Electro, star.FIVE, "大风纪官")

	TravelerDendro = New("Traveler Dendro", elemental.Dendro, star.FIVE, "草主")
	Tighnari       = New("Tighnari", elemental.Dendro, star.FIVE, "驴耳朵")
	Collei         = New("Collei", elemental.Dendro, star.FOUR)
	Nahida         = New("Nahida", elemental.Dendro, star.FIVE, "小草神", "小吉祥草王")
	Alhaitham      = New("Alhaitham", elemental.Dendro, star.FIVE, "大书记官", "海哥")
	Yaoyao         = New("Yaoyao", elemental.Dendro, star.FOUR)

	Chongyun      = New("Chongyun", elemental.Cryo, star.FOUR)
	Qiqi          = New("Qiqi", elemental.Cryo, star.FIVE)
	Kaeya         = New("Kaeya", elemental.Cryo, star.FOUR, "凯子哥")
	Diona         = New("Diona", elemental.Cryo, star.FOUR, "猫猫")
	Ganyu         = New("Ganyu", elemental.Cryo, star.FIVE, "王小美", "椰羊")
	Rosaria       = New("Rosaria", elemental.Cryo, star.FOUR)
	Eula          = New("Eula", elemental.Cryo, star.FIVE)
	KamisatoAyaka = New("Kamisato Ayaka", elemental.Cryo, star.FIVE, "白鹭公主")
	Aloy          = New("Aloy", elemental.Cryo, star.FIVE)
	Shenhe        = New("Shenhe", elemental.Cryo, star.FIVE, "小姨")
	Layla         = New("Layla", elemental.Cryo, star.FOUR)
	Mika          = New("Mika", elemental.Cryo, star.FOUR)

	TravelerGeo = New("Traveler Geo", elemental.Geo, star.FIVE, "岩主")
	Noelle      = New("Noelle", elemental.Geo, star.FOUR, "女仆")
	Ningguang   = New("Ningguang", elemental.Geo, star.FOUR, "富婆", "天权星")
	Zhongli     = New("Zhongli", elemental.Geo, star.FIVE, "岩神", "岩王帝君", "岩王爷", "摩拉克斯")
	Albedo      = New("Albedo", elemental.Geo, star.FIVE, "白垩之子")
	AratakiItto = New("Arataki Itto", elemental.Geo, star.FIVE, "荒泷天下第一斗", "放牛的")
	Gorou       = New("Gorou", elemental.Geo, star.FOUR, "希娜小姐")
	YunJin      = New("Yun Jin", elemental.Geo, star.FOUR, "云先生")
)

var (
	nameMap = make(map[string]*Character)
)

func New(id string, elemental *elemental.Elemental, star int, alias ...string) *Character {
	names := i18n.TS(id)
	names = append(names, alias...)
	c := &Character{
		Id:        id,
		Names:     names,
		Elemental: elemental,
	}
	for _, n := range names {
		nameMap[strings.ToLower(n)] = c
	}
	return c
}

func (c *Character) Name(ctx telebot.Context) string {
	return i18n.T(ctx, c.Id)
}

func Search(name string) []*Character {
	return object.Search(name, nameMap)
}
