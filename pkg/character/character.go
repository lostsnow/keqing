package character

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"

	"gopkg.in/telebot.v3"
	"gopkg.in/yaml.v3"

	"github.com/lostsnow/keqing/data"
	"github.com/lostsnow/keqing/pkg/elemental"
	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
	"github.com/lostsnow/keqing/pkg/object"
	"github.com/lostsnow/keqing/pkg/star"
)

var ErrInvalidElemental = errors.New("invalid elemental")

type Character struct {
	ID          string               `yaml:"id"`
	Star        star.Star            `yaml:"star"`
	Elemental   *elemental.Elemental `yaml:"-"`
	Names       []string             `yaml:"names"`
	NameAliases []string             `yaml:"nameAliases"`
}

func (c *Character) UnmarshalYAML(n *yaml.Node) error {
	type C Character

	type T struct {
		*C        `yaml:",inline"`
		Elemental string `yaml:"elemental"`
	}

	obj := &T{C: (*C)(c)}
	if err := n.Decode(obj); err != nil {
		return err
	}

	elem := elemental.Get(obj.Elemental)

	if elem == nil {
		return fmt.Errorf("%w: %s", ErrInvalidElemental, obj.Elemental)
	}

	c.Elemental = elem

	return nil
}

var (
	objectMap    = make(map[string]*Character)
	nameAliasMap = make(map[string]string)
)

func Init() error {
	return fs.WalkDir(data.Model, "model/character", func(path string, d fs.DirEntry, err error) error {
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
		var v Character
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
}

func (c *Character) Name(ctx telebot.Context) string {
	return i18n.T(ctx, c.ID)
}

func ObjectMap() map[string]*Character {
	return objectMap
}

func Search(name string) []*Character {
	return object.Search(name, nameAliasMap, objectMap)
}
