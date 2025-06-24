package cache_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/dkr290/go-projects/genericcache/internal/cache"
)

func TestCache_Parallel_goroutines(t *testing.T) {
	t.Parallel()
	c := cache.New[int, string]()
	var wg sync.WaitGroup
	const parallelTasks = 10

	wg.Add(parallelTasks)
	for i := range parallelTasks {
		go func(j int) {
			defer wg.Done()
			_ = c.Upsert(4, fmt.Sprint(j))
		}(i)
	}
	wg.Wait()

	wg.Add(parallelTasks)
	for i := range parallelTasks {
		go func(j int) {
			defer wg.Done()
			_ = c.Upsert(i, fmt.Sprintf("key%d", j))
		}(i)
	}
	wg.Wait()

	wg.Add(parallelTasks)
	for i := range parallelTasks {
		go func(j int) {
			defer wg.Done()
			_, _ = c.Read(i)
		}(i)
	}
	wg.Wait()
}
