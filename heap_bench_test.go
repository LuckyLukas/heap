package binaryheap

import (
	"testing"
	"time"
)
import "github.com/valyala/fastrand"

func BenchmarkHeap_AddNoCap(b *testing.B) {
	heap := New(0)
	rands := make(chan *IntNode, 1000000)
	wait := make(chan struct{})
	go func() {
		for {
			select {
			case <- wait:
				return
			default:
				rands <- &IntNode{int64(fastrand.Uint32())}
			}
		}
	}()
	time.Sleep(500 * time.Millisecond) // spin up rng
	b.ResetTimer()
	b.SetParallelism(1)

	for n := 0; n < b.N; n++ {
		heap.Add(<-rands)
	}
	wait <- struct{}{}
}


func BenchmarkHeap_AddOverCapacity(b *testing.B) {
	heap := New(1000000)
	rands := make(chan *IntNode, 1000000)
	wait := make(chan struct{})
	go func() {
		for {
			select {
			case <- wait:
				return
			default:
				rands <- &IntNode{int64(fastrand.Uint32())}
			}
		}
	}()
	time.Sleep(500 * time.Millisecond) // spin up rng
	b.ResetTimer()
	b.SetParallelism(1)

	for n := 0; n < b.N; n++ {
		heap.Add(<-rands)
	}
	wait <- struct{}{}
}


func BenchmarkHeap_RemoveTopGrowingTree(b *testing.B) {
	heap := New(0)
	rands := make(chan *IntNode, 1000000)
	wait := make(chan struct{})
	go func() {
		for {
			select {
			case <- wait:
				return
			default:
				rands <- &IntNode{int64(fastrand.Uint32())}
			}
		}
	}()
	time.Sleep(500 * time.Millisecond) // spin up rng
	b.ResetTimer()
	b.SetParallelism(1)

	for n := 0; n < b.N; n++ {
		//add 2, remove one --> tree is growing
		heap.Add(<-rands)
		heap.Add(<-rands)

		heap.RemoveTop()
	}
	wait <- struct{}{}
}


func BenchmarkHeap_RemoveTopStagnantTree(b *testing.B) {
	heap := New(0)
	rands := make(chan *IntNode, 1000000)
	wait := make(chan struct{})
	go func() {
		for {
			select {
			case <- wait:
				return
			default:
				rands <- &IntNode{int64(fastrand.Uint32())}
			}
		}
	}()
	time.Sleep(500 * time.Millisecond) // spin up rng
	//setup a tree
	for i := 0; i< 100;i++ {
		heap.Add(<-rands)
	}
	b.ResetTimer()
	b.SetParallelism(1)
	for n := 0; n < b.N; n++ {
		//add one remove one --> tree is stagnating in size
		heap.Add(<-rands)
		heap.RemoveTop()
	}
	wait <- struct{}{}
}
