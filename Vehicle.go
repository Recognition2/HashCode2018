package main

import (
	"log"
	"bytes"
	"strconv"
)

type Vehicle struct {
	id               int
	pos              Point
	handledRideIDs   []int
	busy             bool
	unavailableUntil int
}

func (v *Vehicle) handleRide (r Ride, t int) {
	if v.busy {
		log.Fatal("Could not handle ride")
		return
	}

	v.busy = true
	v.unavailableUntil = t +
		v.pos.Dist(r.startPoint) + // Distance from vehicle to start of ride
		r.startPoint.Dist(r.endPoint) // Ride distance
	v.handledRideIDs = append(v.handledRideIDs, r.id)

	v.pos = r.endPoint // After this ride, we'll be at the ride's endpoint.
}

func (v *Vehicle) print(b *bytes.Buffer) {
	// The amount of Rides this Vehicle performed
	l := strconv.Itoa(len(v.handledRideIDs))

	// If the vehicle is still busy, then it hasn't fully completed the last ride.


	b.WriteString(l)

	// For every ride, print it's ID
	for _, r := range v.handledRideIDs {
		s := strconv.Itoa(r)
		b.WriteString(s)
		b.WriteByte(' ')
	}

	b.WriteByte('\n')
}