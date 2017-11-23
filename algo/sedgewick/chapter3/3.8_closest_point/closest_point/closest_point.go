package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type point struct {
	x, y float64
}

func distance(p1, p2 point) float64 {
	dx, dy := p2.x-p1.x, p2.y-p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	var cnt int

	d := 0.5
	n := 100000
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
	var points []point
	{
		start := time.Now()
		points = make([]point, n)
		for i := range points {
			points[i].x = rand.Float64()
			points[i].y = rand.Float64()
		}
		end := time.Now()
		fmt.Println("filling", end.Sub(start))
	}
	{
		start := time.Now()
		for i := range points {
			for j := range points {
				p1, p2 := points[i], points[j]
				dx, dy := p2.x-p1.x, p2.y-p1.y
				distance := math.Sqrt(dx*dx + dy*dy)
				if distance < d {
					cnt++
				}
			}
		}
		end := time.Now()

		fmt.Println(cnt, " edges shorter than ", d)

		fmt.Println(end.Sub(start))
	}
}
