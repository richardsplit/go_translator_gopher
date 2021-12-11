package history

import (
	"sort"
	"strings"
	"sync"
)

var once sync.Once

type history map[string]string

var (
	instance history
)

func History() history {
	once.Do(func() {
		instance = history{}
	})
	return instance
}

func (h history) Add(key, value string) {
	h[key] = value
}

func (h history) GetArranged() []map[string]string {
	keys := make([]string, 0, len(h))

	for k := range h {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

	result := []map[string]string{}
	for _, k := range keys {
		result = append(result, map[string]string{k: h[k]})
	}

	return result
}
