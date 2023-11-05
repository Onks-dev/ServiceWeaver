package main

import (
	"context"
	"github.com/ServiceWeaver/weaver"
)

// Reverser component
type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

// Implementation of the reverser component
type reverser struct {
	weaver.Implements[Reverser]
	cache Cache
}

func (r *reverser) Init(context.Context) error {
	var err error
	r.cache, err = weaver.Get[Cache](r)
	return err
}

func (r reverser) Reverse(ctx context.Context, s string) (string, error) {
	rev, err := r.cache.Get(ctx, s)
	if err == nil {
		return rev, nil
	}
	r.Logger().Error("", err)

	// Number
	addCount.Add(1.0)

	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}
