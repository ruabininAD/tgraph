package testing_test

import (
	"log"
	"os"
	"test/src/pkg/generator"
	"testing"
)

func setLogTest() {
	logFile, err := os.OpenFile("logTest.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
}

func DijkstraWrapper(b *testing.B) {
	setLogTest()
	my, err := generator.GetEndingGraph(10, 10)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		my.Dijkstra(0)
	}
}

func BenchmarkMatrix_Dijkstra(b *testing.B) {
	DijkstraWrapper(b)
}

func BellmanFordWrapper(b *testing.B) {
	setLogTest()
	my, err := generator.GetEndingGraph(10, 10)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		my.BellmanFord(0)
	}
}

func BenchmarkMatrix_BellmanFord(b *testing.B) {
	BellmanFordWrapper(b)
}

// https://github.com/ruabininAD/tgraph.git
// go test -bench. -benchmem  algo_test.go
// число операций, которые удалось выполнить
// функция выполняется N наночекунд - ns/op - наносек на операцию
// функция занимает N байт на операцию B/op байт на операцию
// число алокаций памяти
//  go test -bench. algo_test.go
//

// go test -bench. -benchmem -cpuprofile=".\bench_results\cpu.out" -memprofile=".\bench_results\mem.out" -memprofilerate=1  algo_test.go

//pprof

// go tool pprof testing.test.exe .\bench_results\cpu.out расходы по cpu
// list Dijkstra
