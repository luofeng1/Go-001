package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Bucket ..
type Bucket struct {
	Suc         int64 // 成功次数
	Fail        int64 // 失败次数
	WindowStart int64 // 滑动开始时间(毫秒)
}

// BucketedCounter 提供了基本的桶计数器实现
type BucketedCounter struct {
	mux                sync.RWMutex // 获取某个桶时 - 加锁
	TimeInMilliseconds int64        // 滑动窗口的长度（时间间隔）
	NumBuckets         int64        // 滑动窗口中桶的个数
	Tail               int64        // 最后一位值
	Buckets            []*Bucket    // 一组桶
}

// NewBucketedCounter ..
func NewBucketedCounter(NumBuckets, TimeInMilliseconds int64) *BucketedCounter {
	return &BucketedCounter{
		TimeInMilliseconds: TimeInMilliseconds,
		NumBuckets:         NumBuckets,
		Buckets:            make([]*Bucket, NumBuckets),
		Tail:               0,
	}
}

// getCurrent 获取当前的bucket
func (b *BucketedCounter) getCurrent() *Bucket {
	b.mux.Lock()
	defer b.mux.Unlock()

	current := time.Now().UnixNano() / 1e6
	// 首次为空,初始化
	if b.Tail == 0 && b.Buckets[b.Tail] == nil {
		b.Buckets[b.Tail] = &Bucket{
			Suc:         0,
			Fail:        0,
			WindowStart: current,
		}
		return b.Buckets[b.Tail]
	}

	last := b.Buckets[b.Tail]
	if current < last.WindowStart+b.TimeInMilliseconds {
		return last
	}

	for i := int64(0); i < b.NumBuckets; i++ {
		last := b.Buckets[b.Tail]
		if current < last.WindowStart+b.TimeInMilliseconds {
			return last
		} else if current-(last.WindowStart+b.TimeInMilliseconds) > b.NumBuckets*b.TimeInMilliseconds {
			b.Tail = 0
			b.Buckets = make([]*Bucket, b.NumBuckets)
			b.mux.Unlock()
			return b.getCurrent()
		} else {
			b.Tail++
			bucket := &Bucket{
				Suc:         0,
				Fail:        0,
				WindowStart: last.WindowStart + b.TimeInMilliseconds,
			}
			if b.Tail >= b.NumBuckets {
				copy(b.Buckets[:], b.Buckets[1:])
				b.Tail--
			}

			b.Buckets[b.Tail] = bucket
		}
	}

	return b.Buckets[b.Tail]
}

// AddSuc 增加成功
func (b *BucketedCounter) AddSuc() {
	bucket := b.getCurrent()
	atomic.AddInt64(&bucket.Suc, 1)
}

// AddFail 增加失败
func (b *BucketedCounter) AddFail() {
	bucket := b.getCurrent()
	atomic.AddInt64(&bucket.Fail, 1)
}

// GetSucFail 查询成功,失败数
func (b *BucketedCounter) GetSucFail() (suc int64, fail int64) {
	b.mux.RLock()
	defer b.mux.RUnlock()
	for _, v := range b.Buckets {
		if v == nil {
			break
		}
		suc += v.Suc
		fail += v.Fail
	}
	return
}

func main() {
	b := NewBucketedCounter(10, 200)
	b.AddSuc()
	b.AddFail()
	time.Sleep(300 * time.Millisecond)
	b.AddSuc()
	time.Sleep(600 * time.Millisecond)
	b.AddSuc()
	fmt.Println(b.GetSucFail())
}
