package hzl

import (
	"context"
	"sync"

	"github.com/hazelcast/hazelcast-go-client"
)

type Client struct {
	ctx       context.Context
	err       error
	once      sync.Once
	done      chan struct{}
	config    *hazelcast.Config
	hazelcast *hazelcast.Client
	hzlMap    cache.Map[any]
}

func (c *Client) Close() error {
	c.once.Do(func() {
		if c.hazelcast != nil {
			c.err = c.hazelcast.Shutdown(c.ctx)
		}
		close(c.done)
	})
	return c.err
}

func (c *Client) Done() <-chan struct{} {
	return c.done
}

func (c *Client) Error() error {
	return c.err
}

func (c *Client) GetMap() cache.Map[any] {
	return c.hzlMap
}

func (c *Client) Hazelcast() *hazelcast.Client {
	return c.hazelcast
}

func (c *Client) Running() bool {
	if c.hazelcast != nil {
		return c.hazelcast.Running()
	}
	return false
}
