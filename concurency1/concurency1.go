package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func process(data int) int {
	time.Sleep(time.Second * 2)
	return data * 2
}

//func processData(wg *sync.WaitGroup, result *[]int, data int) {
//	//lock.Lock()
//	defer wg.Done()
//	processedData := process(data)
//	//lock.Lock()
//	*result = append(*result, processedData)
//	//lock.Unlock()
//}

func processData(wg *sync.WaitGroup, result *int, data int) {
	defer wg.Done()
	processedData := process(data)
	*result = processedData
}

func main() {
	//start := time.Now()
	//input := []int{1, 2, 3, 4, 5}
	//resutl := []int{}
	//for _, v := range input {
	//	time.Sleep(2 * time.Second)
	//	resutl = append(resutl, v*2)
	//}
	//
	//fmt.Println(time.Since(start))
	//fmt.Println(resutl)

	start := time.Now()
	var wg sync.WaitGroup
	input := []int{1, 2, 3, 4, 5}
	//result := []int{}
	result := make([]int, len(input))
	for i, data := range input {
		wg.Add(1)
		//go processData(&wg, &result, data)
		go processData(&wg, &result[i], data)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(result)
}
