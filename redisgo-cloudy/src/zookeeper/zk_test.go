package zookeeper

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestZkConnection(t *testing.T) {
	con, err := Connect([]string{"192.168.192.135:2181"}, 30*time.Second)
	fmt.Printf("error 。。。。。:  %v", err)
	//con.Create("/test/factory")
	//con.Create("/test/shard")
	//con.Create("/test/shard/health")
	//con.CreateSeqNode("/test/shard/health", "node1", "191.168.192.134:3196")
	childs, eventChan, err := con.getChildrenWatch("/test/shard/health")
	if err != nil {
		fmt.Println("get childe excption", err.Error())
	}
	for index, value := range childs {
		strdata, err := con.getData("/test/shard/health/" + value)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(strconv.Itoa(index), "child data:", strdata)
	}
	fmt.Println(strings.Join(childs, "|"))
	for {
		event := <-eventChan
		eventType := event.Type
		if eventType == EventNodeChildrenChanged {
			fmt.Printf("get data delete ")
			break
		}
	}
}
