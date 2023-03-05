package weapon

import (
	"fmt"
	"io/fs"
	"strings"

	"gopkg.in/telebot.v3"
	"gopkg.in/yaml.v3"

	"github.com/lostsnow/keqing/data"
	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
	"github.com/lostsnow/keqing/pkg/object"
	"github.com/lostsnow/keqing/pkg/star"
	"github.com/lostsnow/keqing/pkg/weapon/types"
)

type Weapon struct {
	Id          string      `yaml:"id"`
	Star        star.Star   `yaml:"star"`
	Type        *types.Type `yaml:"-"`
	Names       []string    `yaml:"names"`
	NameAliases []string    `yaml:"nameAliases"`
}

func (w *Weapon) UnmarshalYAML(n *yaml.Node) error {
	type W Weapon
	type T struct {
		*W   `yaml:",inline"`
		Type string `yaml:"type"`
	}

	obj := &T{W: (*W)(w)}
	if err := n.Decode(obj); err != nil {
		return err
	}

	elem := types.Get(obj.Type)
	if elem == nil {
		return fmt.Errorf("invalid weapon type: %s", obj.Type)
	}
	w.Type = elem
	return nil
}

func Init() error {
	return fs.WalkDir(data.Model, "model/weapon", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() || !strings.HasSuffix(path, ".yml") {
			return nil
		}
		yamlFile, err := data.Model.ReadFile(path)
		if err != nil {
			return nil
		}
		var v Weapon
		err = yaml.Unmarshal(yamlFile, &v)
		if err != nil {
			return fmt.Errorf("unmarshal model file %s failed: %s", path, err)
		}

		objectMap[v.Id] = &v
		nameAliases := i18n.TS(v.Id)
		nameAliases = append(nameAliases, v.NameAliases...)
		for _, nameAlias := range nameAliases {
			nameAliasMap[nameAlias] = v.Id
		}
		return nil
	})
}

var (
	objectMap    = make(map[string]*Weapon)
	nameAliasMap = make(map[string]string)
)

func (w *Weapon) Name(ctx telebot.Context) string {
	return i18n.T(ctx, w.Id)
}

func ObjectMap() map[string]*Weapon {
	return objectMap
}

func Search(name string) []*Weapon {
	return object.Search(name, nameAliasMap, objectMap)
}
