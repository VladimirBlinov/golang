package main

import (
	"fmt"
	metrconv "golang/metrconv/conv"
	"os"
	"strconv"
)

func main() {
	fmt.Println("args", os.Args)
	for _, arg := range os.Args[1:] {
		fmt.Println("arg ", arg)
		t, _ := strconv.ParseFloat(arg, 64)
		km := metrconv.Kilometr(t)
		fmt.Printf("%s = %s, %s = %s\n", km, metrconv.KmToM(km), km, metrconv.KmToMM(km))
	}
}
