package matrix

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Matrix struct {
	n    int
	data [][]int
}

func NewMatrix(n int) *Matrix {
	m := Matrix{n: n, data: make([][]int, n)}
	for i := 0; i < n; i++ {
		m.data[i] = make([]int, n)
	}
	return &m
}

func (m *Matrix) Set(row, col, value int) {
	m.data[row][col] = value
}

func (m *Matrix) SetUnOrientedE(row, col, value int) {
	m.data[row][col] = value
	m.data[col][row] = value
}

func (m *Matrix) Get(row, col int) int {
	return m.data[row][col]
}

func (m *Matrix) PrintLabel(text string) {
	fmt.Println(text)
	m.Print()
}

func (m *Matrix) Print() {

	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			fmt.Printf("%d, ", m.data[i][j])
			//fmt.Printf("%d ", m.data[i][j])
		}
		fmt.Println()

	}
	fmt.Println()
}

func (m *Matrix) Render() {
	file, err := os.OpenFile("src/graph.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text := ""
	for i := 0; i < m.n; i++ {
		text = ""
		for j := 0; j < m.n; j++ {

			text += fmt.Sprintf("%d, ", m.data[i][j])
		}
		file.WriteString(text + "\n")
	}

}

func (m *Matrix) Weigh() {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			if m.Get(i, j) != 0 {
				m.Set(i, j, rand.Intn(100))
			}
		}
	}
}
