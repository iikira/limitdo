//Package limitdo 控制函数事件调用的次数
package limitdo

import (
	"sync"
	"sync/atomic"
)

type (
	//LimitDo 控制函数调用的次数
	LimitDo struct {
		limit uint64
		done  uint64
		mu    sync.Mutex
	}
)

//New 给定限定次数, 初始化LimitDo
func New(limit uint64) *LimitDo {
	return &LimitDo{limit: limit}
}

//Once 限定1次, 初始化LimitDo
func Once() *LimitDo {
	return &LimitDo{limit: 1}
}

//Twice 限定2次, 初始化LimitDo
func Twice() *LimitDo {
	return &LimitDo{limit: 2}
}

// Do 执行 f, 严格限定调用次数
func (ld *LimitDo) Do(f func()) {
	if atomic.LoadUint64(&ld.done) >= ld.limit {
		return
	}

	ld.mu.Lock()
	defer ld.mu.Unlock()

	if ld.done < ld.limit {
		defer atomic.AddUint64(&ld.done, 1)
		f()
	}
}
