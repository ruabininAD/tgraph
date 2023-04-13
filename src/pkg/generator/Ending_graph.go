package generator

import (
	"log"
	"math"
	"math/rand"
	. "test/src/pkg/matrix"
	"time"
)

func GetEndingGraph(vCount, eCount int) (*Matrix, error) {
	M := NewMatrix(vCount)

	k := 5   // параметр формы распределения
	λ := 2.0 // параметр интенсивности распределения
	res := 0.
	log.Printf("Создаем матрицу смежности с помошью распределения Элдинга\n")
	log.Printf("Распределение Элдинга для параметров\n k = %v - параметр формы распределения \n λ = %.2f -параметр интенсивности распределения \n", k, λ)
	probabilities := make([]float64, 0)
	for i := 0; i < vCount; i++ {
		res = erlangDistribution(float64(i), k, λ)
		probabilities = append(probabilities, res)
		log.Printf(" i =  %v - значение \n вероятность  i = %.2f \n", i, res)
	}

	for i := 0; i < vCount; i++ {
		a := generateRandomNumber(probabilities)
		zeroSlise := make([]int, i+1)
		slise := createRandomSlice(vCount-i-1, a)
		slise = append(zeroSlise, slise...)
		for j, v := range slise {
			M.Set(i, j, v)
		}
	}
	return M, nil

}

// Функция для вычисления плотности вероятности распределения Эрланга
func erlangDistribution(x float64, k int, λ float64) float64 {
	numerator := math.Pow(λ, float64(k)) * math.Pow(x, float64(k-1)) * math.Exp(-λ*x)
	denominator := float64(factorial(k - 1))
	return numerator / denominator
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func generateRandomNumber(probabilities []float64) int {
	sum := 0.0
	for _, p := range probabilities {
		sum += p
	}

	r := rand.Float64() * sum
	for i, p := range probabilities {
		if r < p {
			return i
		}
		r -= p
	}

	return len(probabilities) - 1
}

func createRandomSlice(n, m int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, n) // Создаем срез длины n
	if m > n {
		m = n
	}
	// Заполняем m случайных элементов в срезе
	for i := 0; i < m; i++ {
		index := rand.Intn(n) // Генерируем случайный индекс
		if slice[index] == 1 {
			i--
			continue
		}
		slice[index] = 1 // Записываем значение в срез по случайному индексу
	}

	return slice
}
