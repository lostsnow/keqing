package object

import (
	"strings"
)

type Object[T any] struct {
	ID string
}

func Search[T any](name string, nameAliasMap map[string]string, objectMap map[string]*T) []*T {
	name = strings.ToLower(name)
	uniqMap := make(map[string]struct{})
	objs := make([]*T, 0)
	if nameAliasMap == nil {
		nameAliasMap = make(map[string]string, len(objectMap))
	}
	if len(nameAliasMap) == 0 {
		for id := range objectMap {
			nameAliasMap[id] = id
		}
	}

	for k, v := range nameAliasMap {
		lowerK := strings.ToLower(k)
		if strings.Contains(lowerK, name) {
			if _, ok := uniqMap[v]; ok {
				continue
			}
			uniqMap[v] = struct{}{}
			if lowerK == name {
				objs = append(objs, nil)
				copy(objs[1:], objs)   //nolint:gosec
				objs[0] = objectMap[v] //nolint:gosec
			} else {
				objs = append(objs, objectMap[v])
			}
		}
	}

	return objs
}

func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) { //nolint:nonamedreturns
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}
