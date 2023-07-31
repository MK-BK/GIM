package cache

import (
	"GIM/models"
	"errors"
	"reflect"
	"time"

	goCache "github.com/patrickmn/go-cache"
)

var cache = goCache.New(5*time.Minute, 10*time.Minute)

// Setter --------------------------------------------
func Set(id string, value interface{}) error {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Ptr:
		cache.SetDefault(id, value)
	}

	return errors.New("Value Must Be Ptr")
}

// Getter --------------------------------------------
func GetUser(id string) (*models.User, bool) {
	v, exist := cache.Get(id)
	if !exist {
		return nil, exist
	}

	if user, ok := v.(*models.User); ok {
		return user, exist
	}

	return nil, false
}

func GetSession(id string) (*models.Session, bool) {
	v, exist := cache.Get(id)
	if !exist {
		return nil, exist
	}

	if session, ok := v.(*models.Session); ok {
		return session, exist
	}

	return nil, false
}

func GetGroup(id string) (*models.Group, bool) {
	v, exist := cache.Get(id)
	if !exist {
		return nil, exist
	}

	if group, ok := v.(*models.Group); ok {
		return group, exist
	}

	return nil, false
}
