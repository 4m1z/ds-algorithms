package main

import "math"

func minTimeToVisitAllPoints(points [][]int) (m int) {
	for i := 0; i < len(points); i++ {
		if i+1 >= len(points) {
			break
		}
		c := points[i]
		n := points[i+1]

		cx := c[0]
		cy := c[1]

		nx := n[0]
		ny := n[1]

		dx := math.Abs(float64(nx - cx))
		dy := math.Abs(float64(ny - cy))

		minD := math.Min(dx, dy)
		if minD > 0 {
			m += int(minD)
		}
		m += int(math.Abs(dy - dx))
	}

	return m
}
