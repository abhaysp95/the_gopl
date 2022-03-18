package memo

import "sync"

type Memo struct {
	f Func
	cache map[string]result
	mu sync.RWMutex
}

type Func func(key string) (interface{}, error)

type result struct {
	val interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()  // mu is of sync.Mutex type
	defer memo.mu.Unlock()
	res, ok := memo.cache[key]
	if !ok {
		res.val, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.val, res.err
}

// not shown in book (something I'm testing)
/* func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.RLock()  // acquire shared lock (mu is of sync.RWMutex type)
	res, ok := memo.cache[key]
	if ok {
		memo.mu.RUnlock()
		return res.val, res.err
	}
	memo.mu.RUnlock()

	// acquire exclusive lock
	memo.mu.Lock()
	res.val, res.err = memo.f(key)
	memo.cache[key] = res
	memo.mu.Unlock()  // release the lock
	return res.val, res.err
} */
