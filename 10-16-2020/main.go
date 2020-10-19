package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	c, _ := os.Open(".csv")
	r, _ := csv.NewReader(c).ReadAll()
	n := make([][]int, len(r))
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			if r[i][j] != "" {
				t, _ := strconv.Atoi(r[i][j])
				n[i] = append(n[i], t)
			}
		}
	}
	for i := 0; i < len(n); i++ {
		for h := 0; h < len(n[i]); h++ {
			for k := 0; k < len(n[i])-1; k++ {
				if n[i][k] > n[i][k+1] {
					x := n[i][k]
					n[i][k] = n[i][k+1]
					n[i][k+1] = x
				}
			}
		}
	}
	fmt.Printf("%v\n", n)

}
