package main

import "github.com/ServiceWeaver/weaver/metrics"

var (
	addCount = metrics.NewCounter(
		"count_reverser",
		"The number of times cache didn't return value to a key")

	addGauge = metrics.NewGauge("gauge_concurrent_calls", "Number of goroutines accessing the cache")
)
