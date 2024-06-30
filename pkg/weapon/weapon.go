package weapon

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"

	"gopkg.in/telebot.v3"
	"gopkg.in/yaml.v3"

	"github.com/lostsnow/keqing/data"
	"github.com/lostsnow/keqing/pkg/i18n"
	"github.com/lostsnow/keqing/pkg/object"
	"github.com/lostsnow/keqing/pkg/star"
	"github.com/lostsnow/keqing/pkg/weapon/types"
)

var ErrInvalidWeaponType = errors.New("invalid weapon type")

type Weapon struct {
	ID          string      `yaml:"id"`
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
		return fmt.Errorf("weapon.UnmarshalYAML: %w", err)
	}

	elem := types.Get(obj.Type)
	if elem == nil {
		return fmt.Errorf("%w: %s", ErrInvalidWeaponType, obj.Type)
	}

	w.Type = elem

	return nil
}

func Init() error {
	err := fs.WalkDir(data.Model, "model/weapon", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil //nolint:nilerr
		}

		if d.IsDir() || !strings.HasSuffix(path, ".yml") {
			return nil
		}

		yamlFile, err := data.Model.ReadFile(path)
		if err != nil {
			return nil //nolint:nilerr
		}

		var v Weapon

		err = yaml.Unmarshal(yamlFile, &v)
		if err != nil {
			return fmt.Errorf("unmarshal model file %s failed: %w", path, err)
		}

		objectMap[v.ID] = &v
		nameAliases := i18n.TS(v.ID)
		nameAliases = append(nameAliases, v.NameAliases...)

		for _, nameAlias := range nameAliases {
			nameAliasMap[nameAlias] = v.ID
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("weapon.Init: %w", err)
	}

	return nil
}

var (
	objectMap    = make(map[string]*Weapon)
	nameAliasMap = make(map[string]string)
)

func (w *Weapon) Name(ctx telebot.Context) string {
	return i18n.T(ctx, w.ID)
}

func ObjectMap() map[string]*Weapon {
	return objectMap
}

func Search(name string) []*Weapon {
	return object.Search(name, nameAliasMap, objectMap)
}
