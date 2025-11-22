package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {
	// ваш код здесь
	if size <= 0 {
		return []int{}, fmt.Errorf("len of slice equal <=0")
	}
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	sliceOfInt := make([]int, size)
	for i := 0; i < size; i++ {
		sliceOfInt[i] = rng.Int()
	}

	return sliceOfInt, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	// ваш код здесь
	if len(data) == 0 {
		return 0, fmt.Errorf("len of slice equal <=0")
	}
	if len(data) == 1 {
		return data[0], nil
	}
	max := data[0]
	for j := range data {
		if max < data[j] {
			max = data[j]
		}
	}
	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		fmt.Println("Len of slice equal <=0")
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	max := data[0]
	for j := range data {
		if max < data[j] {
			max = data[j]
		}
	}
	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data, err := generateRandomElements(SIZE)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max, err := maximum(data)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	var wg sync.WaitGroup
	result := make(chan int, CHUNKS)
	stream := CHUNKS
	for i := 0; i < stream; i++ {
		if SIZE <= CHUNKS {
			stream = 1
		}
		sliceStart := i * SIZE / stream
		sliceEnd := sliceStart + SIZE/stream
		if i == CHUNKS-1 {
			sliceEnd = SIZE
		}
		sliceOfData := data[sliceStart:sliceEnd]
		wg.Add(1)
		go func(dataCh []int) {
			defer wg.Done()
			max := maxChunks(dataCh)
			result <- max
		}(sliceOfData)
	}
	wg.Wait()
	close(result)
	max = 0
	for i := range result {
		if i > max {
			max = i
		}
	}
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
