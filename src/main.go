package main

import (
	"fmt"
	"test/src/my_log"
	"test/src/pkg/generator"
)

func main() {
	my_log.SetLoger() // все логи в logfile.txt

	my, err := generator.GetEndingGraph(6, 8)
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	
	my.PrintLabel("this is graph: ")

	my.BellmanFordPrintLabel(1, "Bellman-Ford")
	my.DijkstraPrintLabel(1, "Dijkstra")
	my.Render()

}
