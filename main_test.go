package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	size := []int{10, 5, 12, 8, 0, 4, 7, -3}
	for _, v := range size {
		res, err := generateRandomElements(v)
		if v <= 0 {
			assert.Equal(t, []int{}, res)
			assert.Error(t, err)

		} else {
			assert.Len(t, res, v)
		}
	}

}

func TestMaximum(t *testing.T) {
	data := [][]int{
		{0},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{},
		{5},
		{25, 10, 3, 8, 44, 3, 3, 5, 25, 55},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{-5, -12, -4, -1, -155},
	}
	for i := 0; i < len(data); i++ {
		m, err := maximum(data[i])
		if len(data[i]) == 0 {
			fmt.Printf("Expected: 0 Actual: %d\n", m) // Выводим ожидаемое значение для проверки в консоли (необязательно)
			assert.Equal(t, 0, m)
			assert.Error(t, err)
		}
		if len(data[i]) == 1 {
			fmt.Printf("Expected: %d Actual: %d\n", data[i][0], m) // Выводим ожидаемое значение для проверки в консоли (необязательно)
			assert.Equal(t, data[i][0], m)
		}
		if len(data[i]) > 1 {
			max := data[i][0]
			for j := 0; j < len(data[i]); j++ {
				if data[i][j] > max {
					max = data[i][j]

				}
			}
			fmt.Printf("Expected: %d Actual: %d\n", max, m) // Выводим ожидаемое значение для проверки в консоли (необязательно)
			assert.Equal(t, max, m)
		}
	}
}
