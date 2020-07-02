package hrank

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestRanking(t *testing.T) {
	r := NewRanking()
	count := 100 * 10000
	startTime := time.Now()
	nums := createNums(count)
	for k, v := range nums {
		r.Set(k, v) // 乱序插入
	}
	fmt.Printf("write useTime:%.2fs\n", time.Now().Sub(startTime).Seconds())

	startTime = time.Now()
	for k, v := range nums {
		r.Set(k, v+1) // 乱序更新
	}
	fmt.Printf("reset useTime:%.2fs\n", time.Now().Sub(startTime).Seconds())

	startTime = time.Now()
	for k := range nums {
		should := count - k2v(k)
		if rank := r.Get(k); rank != should {
			t.Errorf("ranking fail %s should %d, but %d", k, should, rank)
		}
	}
	fmt.Printf("read useTime:%.2fs\n", time.Now().Sub(startTime).Seconds())
}

func createNums(count int) map[string]float64 {
	nums := map[string]float64{}
	for i := 0; i < count; i++ {
		nums[strconv.Itoa(i)] = float64(i)
	}
	return nums
}

func k2v(k string) int {
	i, _ := strconv.Atoi(k)
	return i
}
