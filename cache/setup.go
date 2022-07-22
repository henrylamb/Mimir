package cache

import (
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

type Grdb struct {
	GlobalRdb *cache.Cache
}

type Cachy interface {
	SetRedis(key string, obj interface{}) error
	GetRedis(key string) (interface{}, error)
}

func ConnectCache(cacheAddresses map[string]string, storingObject interface{}) Cachy {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: cacheAddresses,
	})

	myCache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	err := myCache.Once(&cache.Item{
		Key:   "username",
		Value: storingObject, // destination
		SetNX: true,
	})
	if err != nil {
		panic(err)
	}
	newCache := Grdb{
		myCache,
	}
	return &newCache
}

//SetRedis to ensure that incorrect data is not grabbed from the cache in favour of the database the key will have a pattern based off the query params of graphql.
func (rdb *Grdb) SetRedis(key string, obj interface{}) error {
	ctx := context.Background()
	newItem := &cache.Item{
		Key:   key,
		Value: obj,
		SetNX: true,
		Ctx:   ctx,
	}
	err := rdb.GlobalRdb.Set(newItem)
	return err
}

//GetRedis this function will look for an item within a redis cache if you input a key and return the item if its there.
func (rdb *Grdb) GetRedis(key string) (interface{}, error) {
	var (
		val interface{}
		err error
	)
	ctx := context.Background()
	err = rdb.GlobalRdb.Get(ctx, key, &val)
	return val, err
}

//new note
