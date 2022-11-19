package main

import (
	"math"
	"runtime"
	"time"
)

// SquareRoot calculation.
func SquareRoot() float64 {
	var x = 0.001
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	return x
}

// SquareRoot of number calculation.
func SquareRootOf(n float64) float64 {
	return math.Sqrt(n)
}

func RunCPUTest(interval, percentage int) {
	logger.
		Info().
		Int("percentage", percentage).
		Int("interval", interval).
		Str("status", "start").
		Msg("Generating CPU Load")

	runMS := 1000 * percentage
	sleepMS := 1000*100 - runMS

	n := runtime.NumCPU()
	quit := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case <-quit:
					return
				default:
					begin := time.Now()
					for {
						if time.Since(begin) > time.Duration(runMS)*time.Microsecond {
							break
						}
					}
					time.Sleep(time.Duration(sleepMS) * time.Microsecond)
				}
			}
		}()
	}

	time.Sleep(time.Duration(interval) * time.Second)
	for i := 0; i < n; i++ {
		quit <- true
	}

	logger.
		Info().
		Str("status", "completed").
		Msg("Generating CPU Load")
}
