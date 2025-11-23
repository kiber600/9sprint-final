package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	size := []int{10, 5, 12, 8, 0, 4, 7, -3, 500, 755, 5000, 125, 324, 748, 10124, 15725, 250531, 1, 155, 324, 747, 89, 66}
	for _, v := range size {
		res := generateRandomElements(v)
		if v <= 0 {
			assert.Equal(t, []int{}, res)
		} else {
			fmt.Printf("Expected: %d Actual: %d len of slice\n", v, len(res)) // Выводим ожидаемое значение для проверки в консоли (необязательно)
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
		{500, 7000, 1555500, 1212121, 32, 158},
	}
	for i := 0; i < len(data); i++ {
		m := maximum(data[i])
		if len(data[i]) == 0 {
			fmt.Printf("Expected: 0 Actual: %d\n", m) // Выводим ожидаемое значение для проверки в консоли (необязательно)
			assert.Equal(t, 0, m)
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
