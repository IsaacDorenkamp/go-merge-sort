package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func merge_sort(array []int) []int {
	output := make([]int, len(array))
	merge_sort_target(array, output)
	return output
}

func merge_sort_target(array []int, destination []int) {
	if len(array) == 1 {
		destination[0] = array[0]
		return
	}

	midpoint := (len(array) / 2)

	slice_a := array[:midpoint]
	slice_b := array[midpoint:]

	var wg sync.WaitGroup
	
	result_a := make([]int, len(slice_a))
	result_b := make([]int, len(slice_b))

	wg.Add(2)
	go merge_sort_routine(slice_a, result_a, &wg)
	go merge_sort_routine(slice_b, result_b, &wg)
	wg.Wait()

	i, j, k := 0, 0, 0
	for k < len(destination) {
		var toAdd int
		switch true {
		case i >= len(result_a):
			toAdd = result_b[j]
			j++
		case j >= len(result_b):
			toAdd = result_a[i]
			i++
		default:
			if result_b[j] <= result_a[i] {
				toAdd = result_b[j]
				j++
			} else {
				toAdd = result_a[i]
				i++
			}
		}

		destination[k] = toAdd
		k++
	}
}

func merge_sort_routine(source []int, output []int, wait_group *sync.WaitGroup) {
	merge_sort_target(source, output)
	wait_group.Done()
}

// Main Function

func main() {
	rand.Seed(time.Now().Unix())

	unsorted := make([]int, 25)
	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = rand.Intn(100)
	}

	fmt.Printf("Unsorted: %v\n", unsorted)
	sorted := merge_sort(unsorted)
	fmt.Printf("Sorted:   %v\n", sorted)
}