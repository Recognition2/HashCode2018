package main

type Point struct {
	row , col int
}



func (p *Point) Dist(o Point) int {
	csum, rsum:= p.col - o.col, p.row - o.row
	if csum < 0 {
		csum = -csum
	}
	if rsum < 0 {
		rsum = -rsum
	}

	return rsum + csum
}
