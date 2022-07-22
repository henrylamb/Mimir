package cache

import (
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	client *redis.Client
)

var (
	key = "key"
	val = "val"
)

func TestConnectCache(t *testing.T) {
	t.Log("The below tests will not actually test the functions that they correlate to. I couldn't figure out how to use another object")
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal("test failed to create mock redis")
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	require.NotNil(t, client, "the cache mock has been created")
}

func TestGrdb_GetRedis(t *testing.T) {
	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(val, nil))

	r := redis.Cmdable(mock)
	res := r.Get(key)
	require.Equal(t, val, res.Val())
}

func TestGrdb_SetRedis(t *testing.T) {
	exp := time.Duration(0)

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, val, exp).Return(redis.NewStatusResult("", nil))

	r := redis.Cmdable(mock)
	err := r.Set(key, val, exp)
	assert.NoError(t, err.Err(), "This should have err being nil")
}
