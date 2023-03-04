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
			if _, ok := objMap[v]; ok {
				continue
			}
			objMap[v] = struct{}{}
			if k == name {
				objs = append(objs, nil)
				copy(objs[1:], objs)
				objs[0] = v
			} else {
				objs = append(objs, v)
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
