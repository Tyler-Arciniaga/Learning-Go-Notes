package main

import (
	"fmt"
	"math"
	"sync"
)

func main(){
	//ex1()
	//ex2()
	ex3()
}

//since we are writing to the same channel from multiple go routines, we use a wait group to ensure the "clean up" goroutine, only runs it's
//internal logic after the other goroutines finish
func ex1(){
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(2)
	go func(){
		defer wg.Done()
		for i := 0; i < 10; i++{
			out <- i
		}
	}()

	go func(){
		defer wg.Done()
		for i := 10; i < 20; i++{
			out <- i
		}
	}()

	go func(){
		wg.Wait()
		close(out)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)

	go func(){
		defer wg2.Done()
		var result []int
		for v := range out{
			result = append(result, v)
		}
		fmt.Println(result)
		fmt.Println("Done!")
	}()

	wg2.Wait()
}

func ex2(){
	c1 := make(chan int)
	c2 := make(chan int)

	go func(){
		defer close(c1)
		for i := 0; i < 10; i++{
			c1 <- i
		}

	}()

	go func(){
		defer close(c2)
		for i := 10; i < 20; i++{
			c2 <- i
		}
	}()


	for {
		select {
		case v1, ok := <- c1:
			if !ok{
				c1 = nil
				break
			}
			fmt.Printf("from chan1: %d\n", v1)
		case v2, ok := <- c2:
			if !ok{
				c2 = nil
				break
			}
			fmt.Printf("from chan2: %d\n", v2)
		}
		if c1 == nil && c2 == nil{
			break
		}
	}
}

func MapInit() map[int]float64{
	m := make(map[int]float64)
	for i := 0 ; i < 100_000; i++{
		m[i] = math.Sqrt(float64(i))
	}
	return m
}
func ex3(){
	onceMap := sync.OnceValue(MapInit)
	var wg sync.WaitGroup
	for i := 0; i < 100_000; i += 1000{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			m := onceMap()
			fmt.Println(m[i])		
		}(i)
	}
	wg.Wait()
}