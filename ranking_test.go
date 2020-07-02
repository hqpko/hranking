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

func TestRanking_Set(t *testing.T) {
	r := NewRanking()
	r.Set("0", 0)
	r.Set("1", 1)
	r.Set("2", 2)

	// reset, 得分相同的情况下，最新更新的 key 排名更靠前
	r.Set("1", 2)
	if r.Get("1") != 1 {
		t.Errorf("ranking set fail")
	}
}

func TestRanking_Copy_Walk(t *testing.T) {
	r := NewRanking()
	count := 10 * 10000
	nums := createNums(count)
	for k, v := range nums {
		r.Set(k, v)
	}

	r2 := r.Copy()
	r2.Walk(func(index int, key string, score float64) {
		if k2v(key) != int(score) {
			t.Errorf("ranking copy fail, key:%s score:%.0f", key, score)
		} else if index != count-int(score) {
			t.Errorf("ranking copy fail, key:%s index:%d", key, index)
		}
	})
}

// map[i]i, ex: 1:1, 2:2
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
