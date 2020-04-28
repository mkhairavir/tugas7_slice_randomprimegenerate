package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sulaisu := make([]int, 0)
	min, max := 15, 100

	for i := 0; i < 30; i++ {
		randomea := rand.Intn(max-min) + min
		fmt.Println("ini randomea", randomea)
		for j := 0; j < 1; j++ {
			tanda := 0
			for k := 1; k <= randomea; k++ {
				if randomea%k == 0 {
					tanda++
				}
			}
			if tanda == 2 && randomea > 1 {
				sulaisu = append(sulaisu, randomea)
			}
		}
	}
	fmt.Println("ini adalah isi prime slice: ", sulaisu)
}
