package object

import (
	"strings"
)

type Object[T any] struct {
	Id string
}

func Search[T any](name string, nameAliasMap map[string]string, objectMap map[string]*T) []*T {
	name = strings.ToLower(name)
	uniqMap := make(map[string]struct{})
	objs := make([]*T, 0)
	for k, v := range nameAliasMap {
		if strings.Contains(k, name) {
			if _, ok := uniqMap[v]; ok {
				continue
			}
			uniqMap[v] = struct{}{}
			if k == name {
				objs = append(objs, nil)
				copy(objs[1:], objs)
				objs[0] = objectMap[v]
			} else {
				objs = append(objs, objectMap[v])
			}
		}
	}
	return objs
}

func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
