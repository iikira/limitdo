package limitdo

import (
	"sync"
	"testing"
)

func TestLimitDo(t *testing.T) {
	l := New(8)
	wg := sync.WaitGroup{}
	c := make(chan struct{}, 100)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		c <- struct{}{}
		go func(i int) {
			defer func() {
				wg.Done()
				<-c
			}()
			l.Do(func() {
				t.Logf("%d\n", i)
			})
		}(i)
	}
	wg.Wait()
}
