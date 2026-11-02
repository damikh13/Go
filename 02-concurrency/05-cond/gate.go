package main

import "sync"

type Gate struct {
	mu     sync.Mutex
	cond   *sync.Cond
	paused bool
}

func NewGate() *Gate {
	g := &Gate{}
	g.cond = sync.NewCond(&g.mu)
	return g
}

func (g *Gate) Pause() {
	g.mu.Lock()
	g.paused = true
	g.mu.Unlock()
}

func (g *Gate) Resume() {
	g.mu.Lock()
	g.paused = false
	g.mu.Unlock()

	g.cond.Broadcast()
}

func (g *Gate) WaitIfPaused() {
	g.mu.Lock()
	for g.paused {
		g.cond.Wait()
	}
	g.mu.Unlock()
}
