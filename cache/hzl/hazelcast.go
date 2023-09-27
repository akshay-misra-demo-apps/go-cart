package hzl

import (
	"context"
	"log"

	"github.com/hazelcast/hazelcast-go-client"
)

type Cache struct {
	Client *hazelcast.Client
}

type Option func(*options)

type options struct {
	// add any configuration options and set these configurations in func withOptions
}

func New(ctx context.Context, config hazelcast.Config, opts ...Option) (*Cache, error) {

	// register a byte array serializer to decode/encode go structs
	// TODO: check if some other serializer make sense for production use case.
	config.Serialization.SetGlobalSerializer(&ByteArraySerializer{})

	c := withOptions(
		&Client{
			ctx:    ctx,
			done:   make(chan struct{}),
			config: &config},
		opts...)

	client, err := hazelcast.StartNewClientWithConfig(context.TODO(), config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	c.hazelcast = client

	go func() {
		<-ctx.Done()
		_ = c.Close()
	}()

	c.hzlMap = NewMap(c)

	return &Cache{
		Client: client,
	}, nil
}

func withOptions(c *Client, opts ...Option) *Client {
	o := &options{}
	for _, op := range opts {
		op(o)
	}

	return c
}
