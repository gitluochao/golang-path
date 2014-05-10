package zookeeper

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/samuel/go-zookeeper/zk"
)

type ZkConn struct {
	conn *zk.Conn
}

var waitGroup sync.WaitGroup
var (
	EventNodeCreated         = zk.EventType(1)
	EventNodeDeleted         = zk.EventType(2)
	EventNodeDataChanged     = zk.EventType(3)
	EventNodeChildrenChanged = zk.EventType(4)
)
var (
	ErrNoChild          = errors.New("zk : node has not have children")
	ErrNoNode           = errors.New("zk: node does not exist")
	ErrSessionExpired   = errors.New("zk: session has been expired by the server")
	ErrConnectionColsed = errors.New("zk: connection closed")
)

func init() {

}
func Connect(servers []string, timeout time.Duration) (*ZkConn, error) {
	conn, session, err := zk.Connect(servers, timeout)
	if err != nil {
		glog.Error("fail Connect to zookeeper: servers" + strings.Join(servers, "") + err.Error())
	}
	fmt.Println("test %v", err)
	go func(sessionChan <-chan zk.Event) {
		for {
			event := <-sessionChan
			state := event.State
			switch state {
			case zk.StateConnected:
				glog.Info("success Connect to zookeeper")
				waitGroup.Done()
			case zk.StateExpired:
				glog.Info("session expired reconnnect")
				Connect(servers, timeout)
			default:
				break
			}
		}

	}(session)
	waitGroup.Add(1)
	waitGroup.Wait()
	zkConn := &ZkConn{conn}
	return zkConn, err
}
func (zkCon *ZkConn) Create(path string) error {
	cpath := ""
	for _, str := range strings.Split(path, "/")[1:] {
		cpath = strings.Join([]string{cpath, str}, "/")
		_, err := zkCon.conn.Create(cpath, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			if err == zk.ErrNodeExists {
				glog.Warning("zk.Create(%s) already exist", cpath)
			} else {
				glog.Error("zk.Create(%s) error: %v ", cpath, err)
				return err
			}
		}
	}
	return nil
}
func (zkCon *ZkConn) CreateEphemeral(root, nodeName, data string) error {
	_, err := zkCon.conn.Create(strings.Join([]string{root, nodeName}, "/"), []byte(data), zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		if err == zk.ErrNodeExists {
			glog.Warning("zk create Ephemeral already exist")
		} else {
			glog.Error("zk.Create Ephemeral Node error")
			return err
		}
	}
	return nil
}
func (zkCon *ZkConn) CreateSeqNode(root, nodeName, data string) error {
	_, err := zkCon.conn.Create(strings.Join([]string{root, nodeName}, "/"), []byte(data), zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		if err == zk.ErrNodeExists {
			glog.Error("zk node alread exist")
		} else {
			glog.Error("zk.Create Persistent Node error" + err.Error())
		}
	}
	return nil
}
func (zkCon *ZkConn) getChildrenWatch(path string) ([]string, <-chan zk.Event, error) {
	childs, stat, eventChan, err := zkCon.conn.ChildrenW(path)
	if err != nil {
		if err == zk.ErrNoNode {
			return nil, nil, ErrNoNode
		}
		if stat == nil {
			return nil, nil, ErrNoNode
		}
		if len(childs) == 0 {
			return nil, nil, ErrNoChild
		}
	}
	return childs, eventChan, err
}

func (zkCon *ZkConn) setData(path, data string, version int32) error {
	stat, err := zkCon.conn.Set(path, []byte(data), version)
	if err != nil {
		if err == zk.ErrNoNode {
			return ErrNoNode
		}
		if stat == nil {
			return ErrNoNode
		}
	}
	return nil
}

func (zkCon *ZkConn) getData(path string) (string, error) {
	data, stat, err := zkCon.conn.Get(path)
	if err != nil {
		if err == zk.ErrNoNode {
			return "", ErrNoNode
		}
		if err == zk.ErrNoNode {
			return "", ErrNoNode
		}
	}
	if stat == nil {
		return "", ErrNoNode
	}
	return string(data), nil
}
func elect(string path) {
	fmt.Println("begin select new master")
}
