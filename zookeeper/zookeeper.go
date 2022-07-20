package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"github.com/smartwalle/nmutex"
	"path"
)

const (
	kPrefix = "/lock/"
)

type session struct {
	key  string
	conn *zk.Conn
	acl  []zk.ACL
}

func NewSession(key string, conn *zk.Conn, acl []zk.ACL) nmutex.Session {
	var s = &session{}
	s.key = key
	s.conn = conn
	s.acl = acl
	return s
}

func (this *session) NewMutex(key string) nmutex.Mutex {
	var nPath = path.Join(kPrefix, this.key, key)
	return zk.NewLock(this.conn, nPath, this.acl)
}
