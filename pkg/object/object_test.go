package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lostsnow/keqing/pkg/object"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	type TestObject struct {
		ID string
	}

	abc := &TestObject{ID: "abc"}
	cde := &TestObject{ID: "cde"}
	def := &TestObject{ID: "def"}

	name := "de"
	objectMap := map[string]*TestObject{
		"abc": abc,
		"cde": cde,
		"def": def,
	}
	nameAliasMap := map[string]string{
		"abc2": "abc",
		"cde2": "cde",
		"def2": "def",
	}

	ms := object.Search(name, nameAliasMap, objectMap)
	assert.NotContains(t, ms, abc)
	assert.Contains(t, ms, def)
	assert.Contains(t, ms, def)
}
