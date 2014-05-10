package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//a shard to redis.pool
type ShardRedis struct {
	Pool       *redis.Pool
	Path       string //zk path
	Healths    []string
	sicks      []string
	RedisCount int
}

//unthread safe
func New(server, password string, poolSize int, idelTime *time.Duration) *ShardRedis {
	shardRedis := new(ShardRedis)
	pool := createPool(server, password, poolSize, idelTime)
	shardRedis.Pool = pool
	//todo redis count
}
func (shardRedis *ShardRedis) updateShard(server, password string, poolSize int, idelTime *time.Duration) {
	if shardRedis.Pool != nil {
		shardRedis.Pool.Close()
	}
	shardRedis.Pool = createPool(server, password, poolSize, idelTime)
	//update node count
}
func createPool(server, password string, poolSize int, idelTime *time.Duration) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     poolSize,
		IdleTimeout: idelTime,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != nil || password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
