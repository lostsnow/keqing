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
	Bennett   = New("Bennett", elemental.Pyro, star.FOUR)
	Klee      = New("Klee", elemental.Pyro, star.FIVE)
	Xiangling = New("Xiangling", elemental.Pyro, star.FOUR)
	Amber     = New("Amber", elemental.Pyro, star.FOUR)
	Diluc     = New("Diluc", elemental.Pyro, star.FIVE, "卢老爷")
	Xinyan    = New("Xinyan", elemental.Pyro, star.FOUR)
	HuTao     = New("Hu Tao", elemental.Pyro, star.FIVE, "胡堂主")
	Yanfei    = New("Yanfei", elemental.Pyro, star.FOUR)
	Yoimiya   = New("Yoimiya", elemental.Pyro, star.FIVE)
	Thoma     = New("Thoma", elemental.Pyro, star.FOUR)
	Dehya     = New("Dehya", elemental.Pyro, star.FIVE)

	Mona              = New("Mona", elemental.Hydro, star.FIVE, "梅姬斯图斯")
	Xingqiu           = New("Xingqiu", elemental.Hydro, star.FOUR)
	Barbara           = New("Barbara", elemental.Hydro, star.FOUR)
	Tartaglia         = New("Tartaglia", elemental.Hydro, star.FIVE, "公子")
	SangonomiyaKokomi = New("Sangonomiya Kokomi", elemental.Hydro, star.FIVE, "观赏鱼")
	KamisatoAyato     = New("Kamisato Ayato", elemental.Hydro, star.FIVE)
	Yelan             = New("Yelan", elemental.Hydro, star.FIVE, "叶天后")
	Candace           = New("Candace", elemental.Hydro, star.FOUR)
	Nilou             = New("Nilou", elemental.Hydro, star.FIVE)

	TravelerAnemo   = New("Traveler Anemo", elemental.Anemo, star.FIVE, "风主")
	Sucrose         = New("Sucrose", elemental.Anemo, star.FOUR)
	Venti           = New("Venti", elemental.Anemo, star.FIVE, "风神")
	Jean            = New("Jean", elemental.Anemo, star.FIVE, "团长")
	Xiao            = New("Xiao", elemental.Anemo, star.FIVE, "三尺五寸仙人")
	KaedeharaKazuha = New("Kaedehara Kazuha", elemental.Anemo, star.FIVE)
	Sayu            = New("Sayu", elemental.Anemo, star.FOUR, "泥头车")
	ShikanoinHeizou = New("Shikanoin Heizou", elemental.Anemo, star.FOUR)
	Faruzan         = New("Faruzan", elemental.Anemo, star.FOUR, "百岁山")
	Wanderer        = New("Wanderer", elemental.Anemo, star.FIVE, "散兵")

	TravelerElectro = New("Traveler Electro", elemental.Electro, star.FIVE, "雷主")
	Keqing          = New("Keqing", elemental.Electro, star.FIVE, "刻师傅", "牛杂师傅")
	Fischl          = New("Fischl", elemental.Electro, star.FOUR, "皇女")
	Beidou          = New("Beidou", elemental.Electro, star.FOUR, "大姐头")
	Razor           = New("Razor", elemental.Electro, star.FOUR)
	Lisa            = New("Lisa", elemental.Electro, star.FOUR)
	KujouSara       = New("Kujou Sara", elemental.Electro, star.FOUR, "九条天狗")
	RaidenShogun    = New("Raiden Shogun", elemental.Electro, star.FIVE, "雷神")
	YaeMiko         = New("Yae Miko", elemental.Electro, star.FIVE, "屑狐狸")
	KukiShinobu     = New("Kuki Shinobu", elemental.Electro, star.FOUR, "97忍")
	Dori            = New("Dori", elemental.Electro, star.FOUR)
	Cyno            = New("Cyno", elemental.Electro, star.FIVE)

	TravelerDendro = New("Traveler Dendro", elemental.Dendro, star.FIVE, "草主")
	Tighnari       = New("Tighnari", elemental.Dendro, star.FIVE)
	Collei         = New("Collei", elemental.Dendro, star.FOUR)
	Nahida         = New("Nahida", elemental.Dendro, star.FIVE, "草神")
	Alhaitham      = New("Alhaitham", elemental.Dendro, star.FIVE)
	Yaoyao         = New("Yaoyao", elemental.Dendro, star.FOUR)

	Chongyun      = New("Chongyun", elemental.Cryo, star.FOUR)
	Qiqi          = New("Qiqi", elemental.Cryo, star.FIVE)
	Kaeya         = New("Kaeya", elemental.Cryo, star.FOUR)
	Diona         = New("Diona", elemental.Cryo, star.FOUR, "猫猫")
	Ganyu         = New("Ganyu", elemental.Cryo, star.FIVE)
	Rosaria       = New("Rosaria", elemental.Cryo, star.FOUR)
	Eula          = New("Eula", elemental.Cryo, star.FIVE)
	KamisatoAyaka = New("Kamisato Ayaka", elemental.Cryo, star.FIVE)
	Aloy          = New("Aloy", elemental.Cryo, star.FIVE)
	Shenhe        = New("Shenhe", elemental.Cryo, star.FIVE, "小姨")
	Layla         = New("Layla", elemental.Cryo, star.FOUR)
	Mika          = New("Mika", elemental.Cryo, star.FOUR)

	TravelerGeo = New("Traveler Geo", elemental.Geo, star.FIVE, "岩主")
	Noelle      = New("Noelle", elemental.Geo, star.FOUR, "女仆")
	Ningguang   = New("Ningguang", elemental.Geo, star.FOUR)
	Zhongli     = New("Zhongli", elemental.Geo, star.FIVE, "岩神", "帝君")
	Albedo      = New("Albedo", elemental.Geo, star.FIVE)
	AratakiItto = New("Arataki Itto", elemental.Geo, star.FIVE)
	Gorou       = New("Gorou", elemental.Geo, star.FOUR, "希娜")
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
