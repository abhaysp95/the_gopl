package main

import (
	"image"
	"sync"
)

var icons = make(map[string]image.Image)

/**
 *
 * func loadIcons() {
 * 	icons = map[string]image.Image{
 * 		"spades.png": loadIcon("spades.png"),
 * 		"hearts.png": loadIcon("hearts.png"),
 * 		"diamonds.png": loadIcon("diamonds.png"),
 * 		"clubs.png": loadIcon("clubs.png"),
 * 	}
 * }
 *
 * In the absence of explicit synchronization, the compiler and CPU are free
 * to reorder accesses in any number of ways, so long as the behavior of each
 * goroutine is sequentially consistent.
 * Below function is just that of what above function is shown in comment
 */
func loadIcons() {
	icons = make(map[string]image.Image)
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamonds.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}

func loadIcon(icon string) image.Image {
	// a demo func which will load the icon file
	return nil
}

// not concurrency-safe
/* func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()  // one-time initialization
	}
	return icons[name]
} */


// var mu sync.Mutex

// concurrency safe
/* func Icon(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
} */

// using multiple-readers lock
/* var mu sync.RWMutex

func Icon(name string) image.Image {
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()

	// acquire exclusive lock
	mu.Lock()
	if icons == nil {
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
} */

// using sync.Once. Conceptually, a Once consists of a mutex and a boolean
// variable that records whether initialization has taken place; the mutex
// guards both the boolean and the client's data structures. The sole method,
// Do, accepts the initialization function as its argument.

var loadIconsOnce sync.Once

// concurrency-safe
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

// Each call to Do(loadIcons) locks the mutex and checks the boolean variable.
// In the first call, in which the variable is false, Do calls loadIcons and
// set the variable to true. Subsequent calls do nothing, but the mutex
// synchronization ensures that the effects of loadIcons on memory
// (specifically icons) becomes visible to all goroutines
