package object

import (
	"strings"
)

type Object[T any] struct {
	Id string
}

func Search[T any](name string, nameMap map[string]*T) []*T {
	name = strings.ToLower(name)
	objMap := make(map[*T]struct{})
	objs := make([]*T, 0)
	for k, v := range nameMap {
		if strings.Contains(k, name) {
			objMap[v] = struct{}{}
			objs = append(objs, v)
		}
	}
	return objs
}
