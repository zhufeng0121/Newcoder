package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var n int64 = 10000000000
var h float64 = 1.0 /float64(n)

func f(a float64) float64 {
	return 4.0 / (1.0 + a * a)
}

func chunk(start, end int64, c chan float64) {
	var sum float64 = 0.0
	for i := start; i < end; i++ {
		x := h * (float64(i) + 0.5)
		sum += f(x)
	}
	c <- sum * h

}

func main() {
	var cpuProfile = flag.String("cpuprofile","","write cpu profile to file")
	var memProfile = flag.String("memprofile","","write mem profile to file")
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	//记录开始时间
	start := time.Now()
	var pi float64

	np := runtime.NumCPU()

	runtime.GOMAXPROCS(np)

	c := make(chan float64, np)

	for i := 0; i < np; i++ {
		//利用多处理器，并发处理
		go chunk(int64(i) * n / int64(np),(int64(i) + 1) * n / int64(np),c)
	}

	for i := 0; i < np; i++ {
		tmp := <- c
		fmt.Println("c->",tmp)

		pi += tmp
		fmt.Println("pai",pi)
	}

	fmt.Println("Pi", pi)

	end := time.Now()

	//输出执行时间，单位为ms
	fmt.Printf("spend time: %vs\n", end.Sub(start).Seconds())

	//采样memory状态
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()

	}



}