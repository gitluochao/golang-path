package heartBeat

import (
	"redis"
	"strings"
	"github.com/garyburd/redigo/redis"
)

//start rediscloudy.conf heartBeat GoRuntimeSize
func monitor(shardMap map[int32]*redis.ShardRedis, inteval, retryTime int) {
	//monitor by group
	for {
		for key, shard := range shardMap {
			shardPool := shard.Pool
			if !sendBeat(shardPool, inteval, retryTime) {
				//change node

			}
		}
	}
}
func sendBeat(shardPool *redis.Pool, inteval, retryTime int) (health bool) {
	health := true
	con := shardPool.Get()
	if con != nil {
		_, perr := con.Do("PING")
		if perr != nil {
			health = false
		}
	}
	for i := 0; i < retryTime; i++ {
		con := shardPool.Get()
		_, err := con.Do("PING")
		if err != null {
			health = false
		} else {
			health = true
		}
	}
	return
}
func checkHealth(shard *redis.ShardRedis) {
	healths := shard.Healths
	sicks := shard.Sicks
	for _,health := healths {
	    //check healths node 
		child = strings.Join([]string{shard.Path,health},"/")
	    
	}
	for _,sick := sicks {
		child = string.Join([])
	}
}
func checkHealth(redis address){
	
}
