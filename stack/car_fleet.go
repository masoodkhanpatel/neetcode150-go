package stack

import "sort"

// Car Fleet
func carFleet(target int, position []int, speed []int) int {
	type car struct {
		pos  int
		time float64
	}
	n := len(position)
	cars := make([]car, n)
	for i := 0; i < n; i++ {
		cars[i] = car{
			pos:  position[i],
			time: float64(target-position[i]) / float64(speed[i]),
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleetCount := 0
	var lastTime float64
	for _, c := range cars {
		if c.time > lastTime {
			fleetCount++
			lastTime = c.time
		}
	}

	return fleetCount
}
