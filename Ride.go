package main

type Ride struct {
	id int
	startPoint Point
	endPoint Point
	startTime int
	finTime int
}

func (r *Ride) rating () int {
	return r.endPoint.Dist(r.startPoint)
}