package data

import (
	"github.com/matthistuff/simplecache"
	"regexp"
)

var cache *simplecache.Cache

func LoadCache() error {
	c, err := simplecache.New(".shelf-cache")

	if err == nil {
		cache = c
	}

	return err
}

func FromCache(key string) (string, bool) {
	val, exists := cache.Get(key)

	if !exists {
		return key, exists
	}

	return val.(string), exists
}

func SetCache(key string, data string) {
	cache.Assert()[key] = data
}

func FlushCache() error {
	return cache.Flush()
}

func ClearCache() {
	cache.Clear()
}

func AssertGuid(maybeGuid string) (val string, exists bool) {
	if isGuid(maybeGuid) {
		return maybeGuid, true
	}

	val, exists = FromCache(maybeGuid)
	return
}

var guidReg, _ = regexp.Compile("[a-f0-9]{24}")

func isGuid(maybeGuid string) bool {
	return guidReg.MatchString(maybeGuid)
}