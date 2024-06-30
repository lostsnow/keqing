package types

import (
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/i18n"
	"github.com/lostsnow/keqing/pkg/object"
)

type Type struct {
	ID    string
	Names []string
}

var (
	Sword    = New("Sword")
	Claymore = New("Claymore")
	Polearm  = New("Polearm")
	Catalyst = New("Catalyst")
	Bow      = New("Bow")
)

var objectMap = make(map[string]*Type)

func New(id string) *Type {
	names := i18n.TS(id)
	e := &Type{
		ID:    id,
		Names: names,
	}

	for _, n := range names {
		objectMap[n] = e
	}

	return e
}

func Get(id string) *Type {
	return objectMap[id]
}

func (e *Type) Name(ctx telebot.Context) string {
	return i18n.T(ctx, e.ID)
}

func Search(name string) []*Type {
	return object.Search(name, nil, objectMap)
}
