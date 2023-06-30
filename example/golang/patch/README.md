## monkey
使用monkey和supermonkey替换golang的函数实现;
```go
package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

// 注意:这个库github.com/cch123/supermonkey直接依赖bou.ke/monkey库,所以也无法在ARM架构的机器上运行
func main() {
	fmt.Println("原始输出:")
	heyHey()

	patchGuard := sm.Patch(heyHey, func() {
		fmt.Println("replace me")
	})
	fmt.Println("替换后输出:")
	heyHey()

	patchGuard.Unpatch()
	fmt.Println("不替换了,重新输出:")
	heyHey()
}

//go:noinline
func heyHey() {
	fmt.Println("hello")
}
```

执行:`go run -gcflags="-l" yourfile.go`

更多使用请[参考](https://github.com/cch123/supermonkey)