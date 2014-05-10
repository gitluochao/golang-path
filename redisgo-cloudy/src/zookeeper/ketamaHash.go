package zookeeper

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

var ErrEmplyCircle = errors.New("empty circle")

type units []uint32

//sort should impl
func (x units) Len() int {
	return len(x)
}

func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type Consistent struct {
	circle     map[uint32]string
	members    map[string]bool
	sortedHash units
	replicas   int
	count      int64
	sync.RWMutex
}

func New(replicas int) *Consistent {
	c := new(Consistent)
	c.replicas = replicas
	c.circle = make(map[uint32]string)
	c.members = make(map[string]bool)
	return c
}

//replicas elm key
func elmKey(key string, i int) string {
	return key + "&" + strconv.Itoa(i)
}

//add Node
func (c *Consistent) Add(key string) {
	c.Lock()
	defer c.Unlock()
	for i := 0; i < c.replicas; i++ {
		ikey := elmKey(key, i)
		c.circle[c.hashKey(ikey)] = key
	}
	c.updateSortedHash()
	c.members[key] = true
	c.count++
}

//remove Node
func (c *Consistent) Remove(key string) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("get ...")
	for i := 0; i < c.replicas; i++ {
		ikey := elmKey(key, i)
		delete(c.circle, c.hashKey(ikey))
	}
	c.updateSortedHash()
	c.members[key] = false
	c.count--
}

//set all node
func (c *Consistent) Set(keys []string) {
	//move deleve node
	c.Lock()
	defer c.Unlock()
	for key := range c.members {
		found := false
		for _, v := range keys {
			if v == key {
				found = true
				continue
			}
		}
		if !found {
			c.Remove(key)
		}
	}
	for _, elmKey := range keys {
		_, exist := c.members[elmKey]
		if exist {
			continue
		}
		c.Add(elmKey)
	}
}
func (c *Consistent) search(key uint32) (i int) {
	f := func(x int) bool {
		return c.sortedHash[x] > key
	}
	i = sort.Search(len(c.sortedHash), f)
	if i >= len(c.sortedHash) {
		return 0
	}
	return
}

//getNode
func (c *Consistent) getNode(key string) (string, error) {
	c.RLock()
	defer c.RUnlock()
	if len(c.circle) == 0 {
		return "", ErrEmplyCircle
	}
	ikey := c.hashKey(key)
	i := c.search(ikey)
	return c.circle[c.sortedHash[i]], nil
}

//sort key
func (c *Consistent) updateSortedHash() {
	sortHashs := c.sortedHash[:0]
	for key := range c.circle {
		sortHashs = append(sortHashs, key)
	}
	sort.Sort(sortHashs)
	c.sortedHash = sortHashs
}
func (c *Consistent) hashKey(key string) uint32 {
	if len(key) < 64 {
		var scrap [64]byte
		copy(scrap[:], key)
		return crc32.ChecksumIEEE(scrap[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}
