package matrix

import (
	"fmt"
	"log"
)

func max(arr []int) int {
	max := arr[0]
	for _, element := range arr {
		if element > max {
			max = element
		}
	}
	return max
}

func min(arr []int) int {
	min := max(arr)
	for _, element := range arr {
		if element < min && element != 0 {
			min = element
		}
	}
	return min
}

func (m *Matrix) Identity() {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			m.Set(i, j, 1)
		}
	}
}

func Add(other, m *Matrix) (*Matrix, error) {
	if m.n != other.n {
		log.Printf("matrices have different sizes: %dx%d and %dx%d", m.n, m.n, other.n, other.n)
		return nil, fmt.Errorf("matrices have different sizes: %dx%d and %dx%d", m.n, m.n, other.n, other.n)
	}

	result := &Matrix{n: m.n, data: make([][]int, m.n)}
	for i := 0; i < m.n; i++ {
		result.data[i] = make([]int, m.n)
		for j := 0; j < m.n; j++ {
			result.data[i][j] = m.data[i][j] + other.data[i][j]
		}
	}

	return result, nil
}

func Multiply(b, a *Matrix) *Matrix {
	if a.n != b.n {
		log.Printf("The matrices cannot be multiplied: n1 =%v, n2 =  %v", a.n, b.n)
		panic("The matrices cannot be multiplied")
	}

	result := NewMatrix(a.n)

	for i := 0; i < a.n; i++ {
		for j := 0; j < a.n; j++ {
			sum := 0
			for k := 0; k < a.n; k++ {
				sum += a.data[i][k] * b.data[k][j]
			}
			result.data[i][j] = sum
		}
	}

	return result
}
