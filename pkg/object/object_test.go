package object

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	type TestObject struct {
		Id string
	}

	abc := &TestObject{Id: "abc"}
	def := &TestObject{Id: "def"}

	name := "abc"
	nameMap := map[string]*TestObject{
		"abc": abc,
		"def": def,
	}

	ms := Search(name, nameMap)
	for _, m := range ms {
		fmt.Println(m.Id)
	}
}
