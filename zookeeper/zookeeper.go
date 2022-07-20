package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"github.com/smartwalle/nmutex"
	"path"
)

const (
	kPrefix = "/mutex/"
)

type session struct {
	domain string
	conn   *zk.Conn
	acl    []zk.ACL
}

func NewSession(domain string, conn *zk.Conn, acl []zk.ACL) nmutex.Session {
	var s = &session{}
	s.domain = domain
	s.conn = conn
	s.acl = acl
	return s
}

func (this *session) NewMutex(key string) nmutex.Mutex {
	var nPath = path.Join("/", this.domain, kPrefix, key)
	return zk.NewLock(this.conn, nPath, this.acl)
}
