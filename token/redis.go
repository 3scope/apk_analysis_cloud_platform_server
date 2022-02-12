package token

import "github.com/gomodule/redigo/redis"

type RedisCache struct {
	ExpireTime   int
	RedisConnect redis.Conn
}

func RedisCacheFactory(expireTime int, host string) (*RedisCache, error) {
	if connect, err := redis.Dial("tcp", host); err != nil {
		return nil, err
	} else {
		return &RedisCache{
			ExpireTime:   expireTime,
			RedisConnect: connect,
		}, nil
	}
}

// To set string value.
func (rc *RedisCache) SetString(key string, imageData string) (err error) {
	_, err = rc.RedisConnect.Do("SETEX", key, rc.ExpireTime, imageData)
	return err
}

// To get string value.
func (rc *RedisCache) GetString(key string) (string, error) {
	var result string
	exists, err := redis.Int(rc.RedisConnect.Do("EXISTS", key))
	if exists > 0 {
		result, err = redis.String(rc.RedisConnect.Do("GET", key))
	}
	return result, err
}

// To delete numberc of keys.
func (rc *RedisCache) DEL(keys ...string) (int, error) {
	var keySlice []interface{}
	for _, key := range keys {
		keySlice = append(keySlice, key)
	}
	return redis.Int(rc.RedisConnect.Do("DEL", keySlice...))
}

func (rc *RedisCache) Close() error {
	return rc.RedisConnect.Close()
}
