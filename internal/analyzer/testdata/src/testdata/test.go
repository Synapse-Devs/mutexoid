package testdata

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
