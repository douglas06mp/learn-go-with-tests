package main

import "sync"

type Counter struct {
	mg    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mg.Lock()
	defer c.mg.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
