package main

import (
	"fmt"
	"sync"
)

const (
	tiny   = 5
	small  = 50
	medium = 100
	large  = 10000
	huge   = 100000
)

func main() {
	sizeArr := []int{tiny, small, medium, large, huge}
	for _, size := range sizeArr {
		count := 0
		wg := sync.WaitGroup{}
		for i := 0; i < size; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				count++
			}()
		}
		wg.Wait()
		fmt.Printf("size:%v  count:%v\n", size, count)
	}

}
