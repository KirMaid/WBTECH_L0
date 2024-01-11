package main

import "sync"

var cache = sync.Map{}

func setCache(key string, value interface{}) {
	cache.Store(key, value)
}

func getCache(key string) (interface{}, bool) {
	return cache.Load(key)
}
