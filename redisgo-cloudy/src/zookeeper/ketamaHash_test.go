package zookeeper

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	consistent := New(1)
	consistent.Add("192.168.192.134:2192")
	consistent.Add("192.168.192.135:2192")
	consistent.Add("192.168.192.135:2190")
	for key, value := range consistent.circle {
		fmt.Println("key:", key, "value:", value)
	}
	fmt.Println("-------------------------")
	for index, val := range consistent.sortedHash {
		fmt.Println("index:", index, "val:", val)
	}
	str, err := consistent.getNode("120003")
	if err != nil {
		fmt.Println("get node exception " + err.Error())
	}
	fmt.Println("node code", str)
	fmt.Println("after delete node ")
	consistent.Remove("192.168.192.135:2190")
	fmt.Println("after delete node ")
	str1, err1 := consistent.getNode("120003")
	if err != nil {
		fmt.Println("get node exception " + err1.Error())
	}
	fmt.Println("node code after delete:", str1)
}
