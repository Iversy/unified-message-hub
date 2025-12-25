package pgstorage

import (
	"fmt"
)

func getShardSchema(sourceChatID uint64) string {
	shardNum := (sourceChatID % 512) + 1
	return fmt.Sprintf("schema_%03d", shardNum)
}

func groupByShard[T any](items []T, getChatID func(T) uint64) map[string][]T {
	groups := make(map[string][]T)
	for _, item := range items {
		shardSchema := getShardSchema(getChatID(item))
		groups[shardSchema] = append(groups[shardSchema], item)
	}
	return groups
}
