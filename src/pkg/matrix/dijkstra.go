package matrix

import (
	"fmt"
	"math"
)

const inf = math.MaxInt32

func (m *Matrix) Dijkstra(start int) (dist []int, prev [][]int) {
	// Количество вершин в графе
	dist = make([]int, m.n)   // Массив минимальных расстояний от начальной вершины
	prev = make([][]int, m.n) // Массив предыдущих вершин для восстановления пути
	for i := 0; i < m.n; i++ {
		dist[i] = inf            // Изначально все расстояния устанавливаем в бесконечность
		prev[i] = make([]int, 0) // Инициализируем пустыми списками
	}
	dist[start] = 0 // Расстояние от начальной вершины до себя равно 0

	visited := make([]bool, m.n) // Массив для отслеживания посещенных вершин

	for count := 0; count < m.n-1; count++ {
		u := -1
		// Выбираем вершину с минимальным расстоянием из еще не посещенных вершин
		for i := 0; i < m.n; i++ {
			if !visited[i] && (u == -1 || dist[i] < dist[u]) {
				u = i
			}
		}

		visited[u] = true // Помечаем вершину как посещенную

		// Обновляем расстояния до соседних вершин через выбранную вершину
		for v := 0; v < m.n; v++ {
			if !visited[v] && m.Get(u, v) != 0 && dist[u] != inf && dist[u]+m.Get(u, v) < dist[v] {
				dist[v] = dist[u] + m.Get(u, v) // Обновляем расстояние
				prev[v] = append(prev[u], u)    // Обновляем путь
			}
		}

	}
	for i, _ := range dist {
		if dist[i] < inf {
			prev[i] = append(prev[i], i)
		}

	}
	return dist, prev
}

func (m *Matrix) DijkstraPrintLabel(start int, label string) {
	fmt.Println(label)
	dist, prev := m.Dijkstra(start)
	for i, v := range prev {
		if dist[i] < inf {
			fmt.Printf("от %v до %v вершины: расстояние =%v,  путь = %v \n", start, i, dist[i], v)
		}

	}
}
