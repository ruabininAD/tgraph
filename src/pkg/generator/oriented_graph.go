package generator

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	. "test/src/pkg/matrix"
	"time"
)

func GetOrientedGraph(vCount, eCount int) (*Matrix, error) {
	M := NewMatrix(vCount)

	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	if eCount > vCount*vCount-vCount {
		log.Printf("graph cannot contain so many arcs:  vCount = %v, eCount =%v", vCount, eCount)
		return nil, errors.New(fmt.Sprintf("graph cannot contain so many arcs:  vCount = %v, eCount =%v", vCount, eCount))
	}

	log.Printf("start generated oriented graph vCount = %v, eCount =%v", vCount, eCount)
	for count := 0; count < eCount; count++ {

		var i int = rand.Intn(vCount)
		var j int = rand.Intn(vCount)
		log.Printf("arc: %v. i: %v, j: %v \n", count, i, j)

		if (M.Get(i, j) != 1) && (i != j) {
			M.Set(i, j, 1)
			log.Printf(" i = %v  j = %v set m[i,j] = 1\n", i, j)
		} else {
			eCount = eCount + 1
		}

	}
	log.Println("oriented graph - OK")
	return M, nil
}
