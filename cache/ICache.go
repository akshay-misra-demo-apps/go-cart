package cache

import (
	"context"
	"time"

	"github.com/hazelcast/hazelcast-go-client"
)

// ICache represents a generic cache interface.
type ICache interface {
	// Put adds a key-value pair to the cache.
	Put(key string, value interface{}) error

	// Get retrieves a value from the cache based on the key.
	Get(key string) (interface{}, error)

	// Remove deletes an entry from the cache based on the key.
	Remove(key string) error

	// Contains checks if a key exists in the cache.
	Contains(key string) (bool, error)
}

type (
	Client interface {
		Close() error
		Done() <-chan struct{}
		Error() error
		GetMap() Map[any]
		GetList() List[any]
		Client() *hazelcast.Client
		Running() bool
	}

	List[T any] interface {
		Add(ctx context.Context, ns, id string, src any) (bool, error)
		Delete(ctx context.Context, ns, id string) error
		GetList(ctx context.Context, ns, id string) ([][]byte, error)
	}

	Map[T any] interface {
		Clear(ctx context.Context, ns string) error
		Delete(ctx context.Context, ns, id string) (int, error)
		Get(ctx context.Context, ns, id string) T
		Insert(ctx context.Context, ns, id string, src T, ttl ...time.Duration) (int, error)
		SetTTL(ctx context.Context, ns, id string, ttl time.Duration) error
		Size(ctx context.Context, ns string) (int, error)
		Update(ctx context.Context, ns, id string, src any, ttl ...time.Duration) (int, error)
	}
)
