package mutexcheck

import "sync"

type BadStruct struct {
	sync.Mutex              // want "mutex should be a pointer type"
	mu         sync.Mutex   // want "mutex should be a pointer type"
	rwmu       sync.RWMutex // want "mutex should be a pointer type"
}

type GoodStruct struct {
	mu   *sync.Mutex
	rwmu *sync.RWMutex
}

type MixedStruct struct {
	mu1  sync.Mutex // want "mutex should be a pointer type"
	mu2  *sync.Mutex
	rwmu sync.RWMutex // want "mutex should be a pointer type"
	rwm2 *sync.RWMutex
}
