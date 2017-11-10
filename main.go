package main

import (
	"fmt"
	"runtime"

	"github.com/SinedVonotirah/gopo/benchs"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	benchs.ORM_MULTI = 1
	benchs.RunBenchmark("gorm")

	benchs.RunBenchmark("beego")

	fmt.Println("\nReports: \n")
	fmt.Print(benchs.MakeReport())
}
