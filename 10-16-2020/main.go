package main

import (
	"os"
	"encoding/csv"
)

func main(){
	c,_ := os.Open(".csv")
	r,_ := csv.NewReader(c).ReadAll()
	n := make([][]int, len(r))
	print(n)
}