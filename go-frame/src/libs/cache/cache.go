package cache

import (
	"fmt"
)

var (
	errNotFound = fmt.Errorf("cache not found")
)

type CacheStorage interface {
	Set(key string, val []byte) error
	Get(key string) ([]byte, error)
}

func CacheNotFound() error {
	return errNotFound
}

func IsCacheNotFound(err error) bool {
	return err == errNotFound
}
