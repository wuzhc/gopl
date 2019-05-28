// 并发非阻塞缓存系统
// 第一获取url之后,会缓存结果
// 第二次获取同个url时,会直接从结果返回
// 对比memo1版本,这个版本主要是面向对象方式去重构代码
// 对比memo2版本,增多互斥锁
// 对比memo3版本,解决并发被阻塞(串行访问)的问题
// 对比memo4版本,解决两次锁的临界区域有多个goroutine访问的问题
// 对比memo5版本,这是一个新的设计,将共享变量的写操作限定在一个goroutine中
package memo6

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct{ requests chan request }

// New returns a memoization of f. Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition. 通知数据已准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition. 等待数据准备完毕
	<-e.ready
	// Send the result to the client. 向客户端发送结果
	response <- e.res
}
