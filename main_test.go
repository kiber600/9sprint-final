package main

import (
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

	expectedData := []int{10, 55, 1, 1555500}
	j := 0
	for i := 0; i < len(data); i++ {
		m := maximum(data[i])
		if len(data[i]) == 0 {
			assert.Equal(t, 0, m)
		}
		if len(data[i]) == 1 {
			assert.Equal(t, data[i][0], m)
		}
		if len(data[i]) > 1 {
			for j < len(expectedData) {
				assert.Equal(t, expectedData[j], m)
				j++
				break
			}

		}
	}
}
