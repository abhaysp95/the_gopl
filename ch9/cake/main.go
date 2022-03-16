/** Even when a variable cannot be confined to a single goroutine for its
* entire lifetime, confinement may still be a solution to the problem of
* concurrent access. For example, it’s common to share a variable between
* goroutines in a pipeline by passing its address from one stage to the next
* over a channel. If each stage of the pipeline refrains from accessing the
* variable after sending it to the next stage , then all accesses to the
* variable are sequential. In effect, the variable is confined to one stage of
* the pipeline, then confined to the next, and so on. This discipline is
* sometimes called "serial confinement" */
package main

type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"  // no need to dereference (this is golang)
		cooked <- cake  // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake  // iced never touches this cake again
	}
}
