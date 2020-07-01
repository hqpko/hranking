# hranking

基于 `Size Balanced Tree` 的排行榜，提供读写 `O(log n)` 的时间复杂度

```go
package main

import (
	"fmt"
	"time"

	hrank "github.com/hqpko/hranking"
)

func main() {
	r := hrank.NewRanking()
	r.Set("user01", 100)
	r.Get("user01") // 排名 1

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

func createNums(count int) map[string]float64 {
	nums := map[string]float64{}
	for i := 0; i < count; i++ {
		nums[fmt.Sprintf("t%d", i)] = float64(i)
	}
	return nums
}

```