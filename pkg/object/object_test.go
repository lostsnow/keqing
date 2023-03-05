package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	type TestObject struct {
		Id string
	}

	abc := &TestObject{Id: "abc"}
	cde := &TestObject{Id: "cde"}
	def := &TestObject{Id: "def"}

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

	ms := Search(name, nameAliasMap, objectMap)
	assert.NotContains(t, ms, abc)
	assert.Contains(t, ms, def)
	assert.Contains(t, ms, def)
}
