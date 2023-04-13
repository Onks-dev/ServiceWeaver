package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"sync"
)

type Cache interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type cache struct {
	weaver.Implements[Cache]
	mu   sync.Mutex
	data map[string]string
}

func (c *cache) Init(context.Context) error {
	c.data = map[string]string{}
	return nil
}

func (c *cache) Set(_ context.Context, key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	return nil
}

func (c *cache) Get(_ context.Context, key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.data[key]
	if !ok {
		return "", fmt.Errorf("key value pair doesn exist for key %s", key)
	}
	return v, nil
}
