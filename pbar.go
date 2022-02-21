package main

import (
	"math"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func mySqrt(x float64) float64 {
	f := float64(x)
	ff := math.Sqrt(f)
	return float64(ff)
}

func main() {
	count := 1000000
	bar := pb.StartNew(count)

	x := float64(0)
	k := float64(5)
	for {
		x = x + 0.001

		// y = 100 (x - k sqrt x) / (x - k ^ 2)
		// Google:
		// https://www.google.com.hk/search?q=y+%3D++%28x+-+sqrt+x%29+%2F+%28x+-+1%29
		y :=((x - k * mySqrt(x) )/ (x - k * k))*float64(count)
		bar.SetCurrent(int64(y))
		time.Sleep(1*time.Millisecond)
	}

	// finish bar
	bar.Finish()
}