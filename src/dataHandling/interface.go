package dataHandling

import (
	"errors"
	"github.com/yaronsumel/geo-coding-test/src/place"
	"sync"
)

// will help us to prevent race conditions
var mtx = sync.Mutex{}

// handlers hold handler in map
var handlers = make(map[string]Interface)

// Interface - data handler interface
type Interface interface {
	// Next - read next place into &p
	// return error at EOF
	Next(p *place.Place) bool
	// Close
	// close all related resources to the handler
	Close() error
}

// RegHandler register handler for use
func RegHandler(handlerName string, handler Interface) {
	mtx.Lock()
	handlers[handlerName] = handler
	mtx.Unlock()
}

// GetHandler return new handler from handlers map
func GetHandler(handlerName string) (Interface, error) {
	mtx.Lock()
	if v, ok := handlers[handlerName]; ok {
		mtx.Unlock()
		return v, nil
	}
	mtx.Unlock()
	return nil, errors.New("Could not find that handler")
}
