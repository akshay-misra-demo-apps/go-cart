package hzl

import (
	"context"
	"time"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/types"
)

type Map struct {
	client Client
}

func NewMap(client Client) *Map {
	return &Map{client: client}
}

func (c *Map) Clear(ctx context.Context, ns string) error {
	return c.getMap(ctx, ns).Clear(ctx)
}

func (c *Map) Delete(ctx context.Context, ns, id string) (s int, err error) {
	r, err := c.getMap(ctx, ns).Remove(ctx, id)
	if err != nil {
		return
	} else if r != nil {
		s.Deleted++
	}
	return
}

func (c *Map) Get(ctx context.Context, ns, id string) db.Result {
	g, err := c.getMap(ctx, ns).Get(ctx, id)
	if err != nil {
		return db.MakeResult(id, nil, err)
	} else if g == nil {
		return db.MakeResult(id, nil, db.ErrNoResult)
	}
	return db.MakeResult(id, g, nil)
}

func (c *Map) GetAll(ctx context.Context, ns string) (*db.ResultSet, error) {
	v, err := c.getMap(ctx, ns).GetValues(ctx)
	if err != nil {
		return nil, err
	}
	return db.NewResultSet(&items{items: v}), nil
}

func (c *Map) GetKey(ctx context.Context, ns, id string, p *path.Path) (zero types.Entry, err error) {
	g, err := c.getMap(ctx, ns).Get(ctx, id)
	if err != nil {
		return
	}
	return shared.NewEntry[[]byte](p.Join("."), g)
}

func (c *Map) GetKeys(ctx context.Context, ns, id string, p []*path.Path) ([]types.Entry, error) {
	g, err := c.getMap(ctx, ns).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return shared.NewEntries[[]byte](path.PathsFrom(p...).Join(), g)
}

func (c *Map) GetView(ctx context.Context, ns, id string) (view *types.SimpleEntryView, err error) {
	return c.getMap(ctx, ns).GetEntryView(ctx, id)
}

func (c *Map) Insert(ctx context.Context, ns, id string, src any, ttl ...time.Duration) (s db.Stat, err error) {
	d, err := db.ReadSrc(src)
	if err != nil {
		return
	}
	return c.put(ctx, c.getMap(ctx, ns), id, d, ttl...)
}

func (c *Map) IsLocked(ctx context.Context, ns string, id string) (bool, error) {
	return c.getMap(ctx, ns).IsLocked(ctx, id)
}

func (c *Map) Lock(ctx context.Context, ns string, id string) error {
	return c.getMap(ctx, ns).Lock(ctx, id)
}

func (c *Map) NewLockContext(ctx context.Context, ns string) context.Context {
	return c.getMap(ctx, ns).NewLockContext(ctx)
}

func (c *Map) SetKey(ctx context.Context, ns, id string, e types.Entry) (err error) {
	m := c.getMap(ctx, ns)
	g, err := m.Get(ctx, id)
	if err != nil {
		return
	}
	b, err := shared.Bytes[[]byte](g)
	if err != nil {
		return
	} else if err = shared.MergeEntry(&e, ".", &b); err != nil {
		return
	}
	return m.Set(ctx, id, b)
}

func (c *Map) SetKeys(ctx context.Context, ns, id string, entries []types.Entry) (err error) {
	m := c.getMap(ctx, ns)
	g, err := m.Get(ctx, id)
	if err != nil {
		return
	}
	b, err := shared.Bytes[[]byte](g)
	if err != nil {
		return
	} else if err = shared.MergeEntries(entries, ".", &b); err != nil {
		return
	}
	return m.Set(ctx, id, b)
}

func (c *Map) SetTTL(ctx context.Context, ns, id string, ttl time.Duration) error {
	return c.getMap(ctx, ns).SetTTL(ctx, id, ttl)
}

func (c *Map) Size(ctx context.Context, ns string) (int, error) {
	return c.getMap(ctx, ns).Size(ctx)
}

func (c *Map) Unlock(ctx context.Context, ns string, id string, force bool) error {
	if force {
		return c.getMap(ctx, ns).ForceUnlock(ctx, id)
	}
	return c.getMap(ctx, ns).Unlock(ctx, id)
}

func (c *Map) Update(ctx context.Context, ns, id string, src any, ttl ...time.Duration) (s db.Stat, err error) {
	d, err := db.ReadSrc(src)
	if err != nil {
		return
	}
	return c.put(ctx, c.getMap(ctx, ns), id, d, ttl...)
}

func (c *Map) getMap(ctx context.Context, name string) *hazelcast.Map {
	m, err := c.client.Hazelcast().GetMap(ctx, name)
	if err != nil {
		panic("hazelcast client GetMap error: " + err.Error())
	}
	return m
}

func (c *Map) put(ctx context.Context, m *hazelcast.Map, id string, src []byte, ttl ...time.Duration) (s db.Stat, err error) {
	s.ID = id
	var put any
	if ttl == nil {
		put, err = m.Put(ctx, id, src)
	} else {
		put, err = m.PutWithTTL(ctx, id, src, fn.Or(ttl...))
	}
	if err != nil {
		return
	} else if put == nil {
		s.Inserted++
	} else {
		s.Updated++
	}
	return
}
