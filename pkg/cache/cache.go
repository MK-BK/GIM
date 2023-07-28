package cache

import (
	"fmt"
	"time"

	goCache "github.com/patrickmn/go-cache"
)

var cache = goCache.New(5*time.Minute, 10*time.Minute)

// Setter --------------------------------------------
func SetUser(id string, user interface{}) error {
	key := fmt.Sprintf("user-%s", id)
	cache.SetDefault(key, user)
	return nil
}

func SetSession(id string, session interface{}) error {
	key := fmt.Sprintf("session-%s", id)
	cache.SetDefault(key, session)
	return nil
}

func SetGroup(id string, group interface{}) error {
	key := fmt.Sprintf("group-%s", id)
	cache.SetDefault(key, group)
	return nil
}

// Getter --------------------------------------------
func GetUser(id string) (interface{}, error) {
	key := fmt.Sprintf("user-%s", id)
	return get(key)
}

func GetSession(id string) (interface{}, error) {
	key := fmt.Sprintf("session-%s", id)
	return get(key)
}

func GetGroup(id string) (interface{}, error) {
	key := fmt.Sprintf("group-%s", id)
	return get(key)
}

func get(key string) (interface{}, error) {
	v, exist := cache.Get(key)
	if !exist {
		return nil, fmt.Errorf("%s, record not found", key)
	}
	return v, nil
}
