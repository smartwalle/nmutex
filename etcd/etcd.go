package etcd

import (
	"context"
	"github.com/smartwalle/nmutex"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"path"
)

const (
	kPrefix = "mutex"
)

type session struct {
	domain string
	client *clientv3.Client
	opts   []concurrency.SessionOption
}

func NewSession(domain string, client *clientv3.Client, opts ...concurrency.SessionOption) nmutex.Session {
	var s = &session{}
	s.domain = domain
	s.client = client
	s.opts = opts
	return s
}

func (this *session) NewMutex(key string) nmutex.Mutex {
	var nPath = path.Join("/", this.domain, kPrefix, key)
	var session, err = concurrency.NewSession(this.client, this.opts...)
	var mu = &mutex{}
	mu.err = err
	mu.session = session
	mu.mu = concurrency.NewMutex(session, nPath)
	return mu
}

type mutex struct {
	err     error
	session *concurrency.Session
	mu      *concurrency.Mutex
}

func (this *mutex) Lock() error {
	if this.err != nil {
		return this.err
	}
	var err = this.mu.Lock(context.TODO())

	if err != nil && this.session != nil {
		this.session.Close()
	}
	return err
}

func (this *mutex) Unlock() error {
	var err error
	if this.mu != nil {
		err = this.mu.Unlock(context.TODO())
	}
	if this.session != nil {
		this.session.Close()
	}
	return err
}
