package memo

import "sync"

type Memo struct {
	f Func
	cache map[string]result
	mu sync.Mutex
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
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		res.val, res.err = memo.f(key)

		// between two critical sections, several goroutines may race to
		// compute f(key) and update the map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.val, res.err
}

/**
* There's still the case when some of the URLs are fetched twice as they (two
* goroutines) almost call f(key) at the same time and cache is overridden again
* by the result. This is called "duplicate supression". */
