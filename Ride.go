package main

import "sort"

type Ride struct {
	id int
	startPoint, endPoint Point
	startTime, finTime int
	handled bool
}

func (r *Ride) rating () int {
	return r.endPoint.Dist(r.startPoint)
}

// Type of a "less" function, defines ordering
type By func(r1, r2 *Ride) bool


func (by By) Sort(rides []Ride) {
	rs := &rideSorter{
		rides:rides,
		by:by,
	}
	sort.Sort(rs)
}

type rideSorter struct {
	rides []Ride
	by func(r1, r2 *Ride) bool // Closure used in
}

func (s *rideSorter) Len() int {
	return len(s.rides)
}
func (s *rideSorter) Swap(i, j int) {
	s.rides[i], s.rides[j] = s.rides[j], s.rides[i]
}

func (s *rideSorter) Less(i, j int) bool {
	return s.by(&s.rides[i], &s.rides[j])
}

