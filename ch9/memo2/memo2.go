// 并发非阻塞缓存系统
// 第一获取url之后,会缓存结果
// 第二次获取同个url时,会直接从结果返回
// 对比memo1版本,这个版本主要是面向对象方式去重构代码
package memo2

// 定义一个memo对象
type Memo struct {
	// 处理方式
	f Func
	// 缓存结果
	cache map[string]result
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
	if v, ok := m.cache[url]; ok {
		return v.value, v.err
	}

	resp, err := m.f(url)
	m.cache[url] = result{resp, err}
	return resp, err
}
