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
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	sliceOfInt := make([]int, size)
	for i := 0; i < size; i++ {
		sliceOfInt[i] = rng.Int()
	}

	return sliceOfInt
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
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

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	var wg sync.WaitGroup
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	result := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {

		sliceStart := i * len(data) / CHUNKS
		sliceEnd := sliceStart + len(data)/CHUNKS
		if i == CHUNKS-1 {
			sliceEnd = len(data)
		}
		sliceOfData := data[sliceStart:sliceEnd]
		wg.Add(1)
		go func(idx int, dataCh []int) {
			defer wg.Done()
			result[idx] = maximum(dataCh)

		}(i, sliceOfData)

	}
	wg.Wait()
	return maximum(result)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
