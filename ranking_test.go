package hrank

import (
	"fmt"
	"testing"
	"time"
)

func TestRanking(t *testing.T) {
	r := NewRanking()
	count := 10000 * 100
	startTime := time.Now()
	nums := map[string]float64{}
	for i := 0; i < count; i++ {
		nums[fmt.Sprintf("t%d", i)] = float64(i)
	}
	for k, v := range nums {
		r.Set(k, v) // 随机设置，不按顺序插入
	}
	fmt.Printf("write useTime:%.2fs\n", time.Now().Sub(startTime).Seconds())

	startTime = time.Now()
	for k, v := range nums {
		if rank := r.Get(k); rank != int(v)+1 {
			t.Errorf("ranking fail %s should %d, but %d", k, int(v)+1, rank)
		}
	}
	fmt.Printf("read useTime:%.2fs\n", time.Now().Sub(startTime).Seconds())
}
