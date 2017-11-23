package main

import (
	"fmt"
	"math"
)

func distance(dx, dy float64) float64 {
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	var cnt float64
	/*
		f, err := os.Create("cpuprofile.out")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	*/
	for j := 0; j < 100000*100000; j++ {
		cnt += distance(float64(j), float64(j+1))
	}
	fmt.Println(cnt)
}
