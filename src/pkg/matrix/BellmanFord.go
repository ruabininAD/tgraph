package matrix

import (
	"fmt"
	"log"
)

func (m *Matrix) BellmanFord(start int) ([]int, []int) {

	dist := make([]int, m.n)
	prev := make([]int, m.n)
	for i := range dist {
		dist[i] = inf // Инициализация расстояний бесконечными значениями
		prev[i] = -1
	}
	dist[start] = 0 // Расстояние от начальной вершины до самой себя равно 0

	for i := 0; i < m.n-1; i++ { // Цикл релаксации |V| - 1 раз
		for u := 0; u < m.n; u++ {
			for v := 0; v < m.n; v++ {
				if m.Get(u, v) != 0 { // Проверка наличия ребра между вершинами
					if dist[u]+m.Get(u, v) < dist[v] { // Релаксация ребра
						dist[v] = dist[u] + m.Get(u, v)
						prev[v] = u
					}
				}
			}
		}
	}

	// Проверка наличия отрицательных циклов
	for u := 0; u < m.n; u++ {
		for v := 0; v < m.n; v++ {
			if m.Get(u, v) != 0 { // Проверка наличия ребра между вершинами
				if dist[u]+m.Get(u, v) < dist[v] {
					log.Printf("в графе содержится отрицательный цикл")
					return nil, nil
				}
			}
		}
	}

	return dist, prev
}

func (m *Matrix) BellmanFordPrintLabel(start int, text string) {
	dist, prev := m.BellmanFord(start)
	fmt.Println(text)
	fmt.Printf("Расстояния от вершины %d до всех остальных вершин:\n", start)
	for i, d := range dist {
		if d < inf {
			fmt.Printf("Вершина %d: %d. путь: %v\n", i, d, prev[i])
		}

	}
}
