package module

import (
	"math"
	"sync"
)

// SNGenertor 序列号生成器的接口类型
type SNGenertor interface {
	// Start 用于获取预设的最小序列号
	Start() uint64
	// Max 用于获取预设的最大序列号
	Max() uint64
	// Next 用于获取下一个序列号
	Next() uint64
	// CycleCount 用于获取循环计数
	CycleCount() uint64
	// Get 用于获得一个序列号并准备下一个序列号
	Get() uint64
}

// NewSNGenertor 会创建一个序列号生成器
// start用于指定第一个序列号的值
// max用于指定序列号的最大值
func NewSNGenertor(start uint64, max uint64) SNGenertor {
	if max == 0 {
		max = math.MaxUint64
	}
	return &mySNGenertor{
		start: start,
		max:   max,
		next:  start,
	}
}

// mySNGenertor 代表序列号生成器的实现类型
type mySNGenertor struct {
	// start 序列号的最小值
	start uint64
	// max 序列号的最大值
	max uint64
	// next 下一个序列号
	next uint64
	// cycleCount 循环的计数
	cycleCount uint64
	// lock 读写锁
	lock sync.RWMutex
}

func (g *mySNGenertor) Start() uint64 {
	return g.start
}

func (g *mySNGenertor) Max() uint64 {
	return g.max
}

func (g *mySNGenertor) Next() uint64 {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.next
}

func (g *mySNGenertor) CycleCount() uint64 {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.cycleCount
}

func (g *mySNGenertor) Get() uint64 {
	g.lock.Lock()
	defer g.lock.Unlock()
	id := g.next
	if id == g.max {
		g.next = g.start
		g.cycleCount++
	} else {
		g.next++
	}
	return id
}
