package main

import "fmt"

type evictionAlgo interface {
	evict(c *cache)
}

type fifo struct{}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type lru struct{}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

type lfu struct{}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}
