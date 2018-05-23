package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//使用多个CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	//使用sync同步包实现，完成后才结束主线程
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}
