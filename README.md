# hranking

基于 `Size Balanced Tree` 的排行榜，提供读写 `O(log n)` 的时间复杂度

> 基本规则是：由大到小排序，且得分相同的情况下，最新更新的 key 排名更靠前
> key,score 的数据格式根据不同需求做相应修改，现有格式为 key:int64, score:int64

#### 简单示例
```go
package main

import (
	"fmt"
	"time"

	hranking "github.com/hqpko/hranking"
)

func main() {
	r := hranking.NewRanking()
	r.Set(1, 100)
	r.Get(2) // 排名 1

	// 百万级插入&读取
	count := 10000 * 100
	nums := createNums(count)
	startTime := time.Now()
	for k, v := range nums {
		r.Set(k, v) // 乱序写入
	}
	fmt.Printf("write useTime:%.2fs\n", time.Now().Sub(startTime).Seconds())

	startTime = time.Now()
	for k := range nums {
		r.Get(k) // 乱序读取排名
	}
	fmt.Printf("read useTime:%.2fs\n", time.Now().Sub(startTime).Seconds())
}

func createNums(count int) map[int64]int64 {
	nums := map[int64]int64{}
	for i := 0; i < count; i++ {
		nums[int64(i)] = int64(i)
	}
	return nums
}

```

#### 功能示例
```go
package main

import (
	"fmt"
	"strconv"

	hranking "github.com/hqpko/hranking"
)

func main() {
	r := hranking.NewRanking()
	for i := 0; i < 10; i++ {
		r.Set(int64(i), int64(i))
	}

	fmt.Println(r.Len()) // 10

	// key 排行位置
	fmt.Println(r.Get(8)) // 2

	// 第 N 名
	key, score := r.GetN(3)
	fmt.Println(key, score) // 7 7

	// 区间
	keys, scores := r.GetRange(4, 8)
	fmt.Println(keys, scores) // [6 5 4 3 2] [6 5 4 3 2]

	// Copy
	r2 := r.Copy()

	// Walk
	r2.Walk(func(index int, key, score int64) {
		fmt.Printf("walk index:%d, key:%d, score:%d\n", index, key, score) // 9->0
	})
}

```