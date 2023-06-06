package statistics

import "sync"


type Stats struct {
	stat uint64
	mx sync.Mutex
}

type StatMap struct {
	m map[string]uint64
	mx sync.Mutex
}

func (s *Stats) Set() {
	s.mx.Lock()
	s.stat++
	s.mx.Unlock()
}

func (s *Stats) Get() uint64{
	return s.stat
}


func (sm *StatMap) Init() {
	sm.m = make(map[string]uint64)
}

func (sm *StatMap) Set(url string) {
	sm.mx.Lock()
	sm.m[url]++
	sm.mx.Unlock()
}

func (sm *StatMap) Get(url string) uint64{
	return sm.m[url]
}