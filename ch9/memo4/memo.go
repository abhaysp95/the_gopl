package memo

import "sync"

type entry struct {
	res result
	ready chan struct{}
}

type Memo struct {
	f Func
	cache map[string]*entry
	mu sync.Mutex
}

type Func func(key string) (interface{}, error)

type result struct {
	val interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()  // acquire the token
	ent := memo.cache[key]
	if ent == nil {
		ent = &entry{ready: make(chan struct{})}
		// this is the first request for the key
		// this goroutine is responsible for getting the result and broadcasting
		// the ready condition
		memo.cache[key] = ent
		memo.mu.Unlock()

		ent.res.val, ent.res.err = memo.f(key)
		close(ent.ready)  // broadcast the ready condition
	} else {
		memo.mu.Unlock()
		<-ent.ready  // wait for ready condition
	}
	return ent.res.val, ent.res.err
}
