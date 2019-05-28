// 并发非阻塞缓存系统
// 第一获取url之后,会缓存结果
// 第二次获取同个url时,会直接从结果返回
// 对比memo1版本,这个版本主要是面向对象方式去重构代码
// 对比memo2版本,增多互斥锁
// 对比memo3版本,解决并发被阻塞(串行访问)的问题
package memo4

import (
	"sync"
)

// 定义一个memo对象
type Memo struct {
	// 处理方式
	f Func
	// 缓存结果
	cache map[string]result
	// 增加互斥锁
	mu sync.Mutex
}

type Func func(string) (interface{}, error)

// 结果也处理为一个对象,因为要返回这个对象
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (m *Memo) Get(url string) (interface{}, error) {
	m.mu.Lock()
	v, ok := m.cache[url]
	m.mu.Unlock()
	if ok {
		return v.value, v.err
	}

	// 在两个临界区域之前,可能会有多个goroutine来计算f(),并且更新map
	resp, err := m.f(url)
	m.mu.Lock()
	m.cache[url] = result{resp, err}
	m.mu.Unlock()
	return resp, err
}
