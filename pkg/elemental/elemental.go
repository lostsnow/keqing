package elemental

import (
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
	"github.com/lostsnow/keqing/pkg/object"
)

type Elemental struct {
	Id    string
	Names []string
}

var (
	Pyro    = New("Pyro")
	Hydro   = New("Hydro")
	Anemo   = New("Anemo")
	Electro = New("Electro")
	Dendro  = New("Dendro")
	Cryo    = New("Cryo")
	Geo     = New("Geo")
)

var (
	objectMap = make(map[string]*Elemental)
)

func New(id string) *Elemental {
	names := i18n.TS(id)
	e := &Elemental{
		Id:    id,
		Names: names,
	}
	for _, n := range names {
		objectMap[n] = e
	}
	return e
}

func Get(id string) *Elemental {
	return objectMap[id]
}

func (e *Elemental) Name(ctx telebot.Context) string {
	return i18n.T(ctx, e.Id)
}

func Search(name string) []*Elemental {
	return object.Search(name, nil, objectMap)
}
