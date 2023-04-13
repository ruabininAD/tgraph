package matrix

import (
	"fmt"
	"log"
)

func (m *Matrix) Shimbel_step(step int, fun string) *Matrix {
	//step соответствует степени
	res := m
	for i := 0; i < step-1; i++ {
		res = ShimbelMultiply(res, m, fun)
	}
	return res
}

func (m *Matrix) ShimbelDistanceMatrix(fun string) *Matrix {
	log.Printf("start ShimbelDistanceMatrix  fun =%s", fun)
	// min  минимальный маршрут от точки до точки
	// max  максимальный маршрут от точки до точки
	res := NewMatrix(m.n)
	ShimbelSteps := make([]*Matrix, 0)
	for i := 1; i < m.n; i++ {
		tmp := m.Shimbel_step(i, fun)
		ShimbelSteps = append(ShimbelSteps, tmp)
	}

	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			arrIJ := make([]int, m.n)
			log.Printf("start generated tab i = %v, j =%v", i, j)
			for step := 0; step < m.n-1; step++ {
				log.Printf("append(arrIJ, ShimbelSteps[step].Get(i, j)) i = %v, j =%v", i, j)
				arrIJ = append(arrIJ, ShimbelSteps[step].Get(i, j))
			}

			value := 0
			if fun == "max" {
				value = max(arrIJ)
			} else {
				value = min(arrIJ)
			}

			res.Set(i, j, value)
		}
	}
	return res
}

func ShimbelMultiply(b, a *Matrix, fun string) *Matrix {
	if a.n != b.n {
		log.Printf("The matrices cannot be multiplied: n1 =%v, n2 =  %v", a.n, b.n)
		panic("The matrices cannot be multiplied")
	}

	result := NewMatrix(a.n)

	for i := 0; i < a.n; i++ {
		for j := 0; j < a.n; j++ {
			arr := make([]int, a.n) //тут массив
			for k := 0; k < a.n; k++ {
				if a.data[i][k] == 0 || b.data[k][j] == 0 {
					arr = append(arr, 0)
					continue
				}
				arr = append(arr, a.data[i][k]+b.data[k][j])
				// добавление в массив
			}
			res := 0
			if fun == "max" {
				res = max(arr)
			} else {
				res = min(arr)
			}

			result.data[i][j] = res // выбор наибольшего из массива.
		}
	}

	return result
}

func (m *Matrix) HowManuRoads(start, stop int) string {
	log.Printf("start HowManuRoads\n")

	if start == stop {
		return fmt.Sprintf("%v is %v:   Minimum road = %v \n", start, stop, 0)
	}

	ShimbelSteps := make([]*Matrix, 0)
	for pow := 1; pow < m.n; pow++ {
		tmp := m.Shimbel_step(pow, "")
		ShimbelSteps = append(ShimbelSteps, tmp)
	}

	roads := make([]int, 0)
	for _, step := range ShimbelSteps {
		roads = append(roads, step.Get(start, stop))
	}

	countRoad := countNZero(roads)
	if countRoad == 0 {
		return "no roads\n"
	} else {
		//return fmt.Sprintf("From %v to %v: There are %v roads. Maximum road = %v, Minimum road = %v \n", start, stop, countRoad, max(roads), min(roads))
		return fmt.Sprintf("From %v to %v: There are %v roads.  Minimum road = %v \n", start, stop, countRoad-1, min(roads))
	}

}

func countNZero(arrRoad []int) (countNZero int) {
	count := 0
	log.Printf("проверка на нули массива: %q\n", fmt.Sprint(arrRoad))
	for lenRoad, countRoad := range arrRoad {
		if lenRoad != 0 {
			log.Printf("имеется %v дорог длинной в %v\n", countRoad, lenRoad)
			count += 1
		}
	}
	return count
}
