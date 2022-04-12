package main

func Gps(s int, x []float64) int {
	maxSpeed := float64(0)
	distance := float64(0)
	for i := range x {
		speed := float64(3600) * (x[i] - distance) / float64(s)
		distance = x[i]
		if speed > maxSpeed {
			maxSpeed = speed
		}
	}
	return int(maxSpeed)
}
