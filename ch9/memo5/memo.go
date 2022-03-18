package memo

// entry contains the result and its ready condition (is broadcasted)
type entry struct {
	res result
	ready chan struct{}
}

// a result is the result of calling func
type result struct {
	val interface{}
	err error
}

// Func is the type of function to memoize
type Func func(key string) (interface{}, error)

type request struct {
	key string
	response chan<- result  // the client wants a single result
}

type Memo struct {
	requests chan request
}

// New returns memoization of f. Client must subsequently call close
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.val, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)  // cache variable is confined to a "monitor goroutine"
	for req := range memo.requests {
		ent := cache[req.key]
		if ent == nil {
			// this is the first request of this key
			ent = &entry{ready: make(chan struct{})}
			cache[req.key] = ent
			go ent.call(f, req.key)
		}
		go ent.deliever(req.response)
	}
}

func (ent *entry) call(f Func, key string) {
	// evaluate the function
	ent.res.val, ent.res.err = f(key)
	// broadcast the ready condition
	close(ent.ready)
}

func (ent *entry) deliever(resp chan<- result) {
	// wait for ready condition
	<-ent.ready
	// send the result to the client
	resp <- ent.res
}
