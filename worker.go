package workers

// WG is Worker group
type WG struct {
	allDone chan bool
	main    chan func()
}

// New creates a new WG
func New(n int) WG {

	procDone := make(chan bool)
	res := WG{
		allDone: make(chan bool),
		main:    make(chan func()),
	}

	for i := 0; i < n; i++ {
		go func() {
			for f := range res.main {
				f()
			}
			procDone <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			_ = <-procDone
		}
		res.allDone <- true

	}()
	return res

}

// Add adds a worker func
func (wg WG) Add(f func()) {
	wg.main <- f
}

// Wait Call Once when all funcs have been added to wait for completion.
// If you don't care to wait, 'go Wait' lol
func (wg WG) Wait() {
	close(wg.main)
	_ = <-wg.allDone

}
