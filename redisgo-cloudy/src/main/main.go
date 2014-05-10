package main

import "fmt"
import (
	"time"
	"github.com/garyburd/redigo/redis"
)

func main() {
	fmt.Printf("Hello world!")
	fmt.Printf("get go ogogog ")
	conn, err := redis.DialTimeout("tcp", "192.168.192.135:6379", 3*time.Second, 3*time.Second, 3*time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	conn.Do("set", "foo", "bar2")
	conn.Do("get", "foo")

}
