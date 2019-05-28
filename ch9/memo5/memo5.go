// 并发非阻塞缓存系统
// 第一获取url之后,会缓存结果
// 第二次获取同个url时,会直接从结果返回
// 对比memo1版本,这个版本主要是面向对象方式去重构代码
// 对比memo2版本,增多互斥锁
// 对比memo3版本,解决并发被阻塞(串行访问)的问题
// 对比memo4版本,解决两次锁的临界区域有多个goroutine访问的问题
package memo5

import (
	"sync"
)

// 定义一个memo对象
type Memo struct {
	// 处理方式
	f Func
	// 缓存结果
	cache map[string]*entry
	// 增加互斥锁
	mu sync.Mutex
}

type Func func(string) (interface{}, error)

// 结果也处理为一个对象,因为要返回这个对象
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res result
	ch  chan struct{}
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (m *Memo) Get(url string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[url]
	if e == nil {
		e = &entry{ch: make(chan struct{})}
		m.cache[url] = e
		m.mu.Unlock()

		e.res.value, e.res.err = m.f(url)
		close(e.ch)
	} else {
		m.mu.Unlock() // 只是更快的释放锁
		<-e.ch
	}

	return e.res.value, e.res.err
}
