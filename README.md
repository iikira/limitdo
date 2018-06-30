# limitdo
控制函数事件调用的次数

# 示例
```go
package main

import (
	"fmt"
	"github.com/iikira/limitdo"
)

func main() {
	ld := limitdo.New(8)
	for i := 0; i < 100; i++ {
		go func(i int) {
			ld.Do(func() {
				fmt.Printf("%d\n", i)
			})
		}(i)
	}
}

```