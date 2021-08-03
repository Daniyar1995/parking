package main

import "fmt"

type car struct {
	start float64
	end   float64
}

type collection struct {
	intervals []car
}

func (i car) inside(start float64) bool {
	if i.start <= start && start < i.end {
		return true
	}
	return false
}

func (c *collection) maxCount() int {
	l := len(c.intervals)
	max := 0

	for i := 0; i < l; i++ {
		start := c.intervals[i].start
		current := 0
		for j := 0; j < l; j++ {
			if c.intervals[j].inside(start) {
				current++
			}
		}

		if current > max {
			max = current
		}
	}
	return max
}

func maxInPark(slots []car) int {
	collection := &collection{slots}
	return collection.maxCount()
}

func main() {
	a := maxInPark([]car{
		car{start: 10, end: 15},
		car{start: 10.40, end: 12},
		car{start: 10, end: 13},
		car{start: 10.50, end: 17},
		car{start: 10, end: 24},
		car{start: 14.18, end: 20},
		car{start: 14, end: 17},
		car{start: 14, end: 17},
		car{start: 14, end: 19},
		car{start: 16, end: 16.20},
		car{start: 17.30, end: 20},
		car{start: 20, end: 22},
	})

	fmt.Println(a)
}
