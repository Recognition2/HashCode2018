package main

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"bytes"
	"io/ioutil"
)

var (
	logE = log.New(os.Stderr, "[ERRO] ", log.Ldate+log.Ltime+log.Ltime+log.Lshortfile)
	logI = log.New(os.Stderr, "[INFO] ", log.Ldate+log.Ltime)
)

type Problem struct {
	Rows int
	Cols int
	Fleet int
	NumRides int
	Bonus int
	TimeSlots int
}


func main () {
	// Start 5 threads, for every example

	file, err := os.Open("example/a_example.in")
	if err != nil {
		logE.Println("Could not open file, quitting")
		os.Exit(1)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	s.Scan()

	p := toProblem(s.Text())
	rides := make([]Ride, p.NumRides)
	vehicles := make([]Vehicle, p.Fleet)

	for i := 0; i < p.NumRides; i++ { // Fill rides with parsed information
		s.Scan()
		rides[i] = toRide(s.Text(), i)
		println("new ride")
	}
	// Create all vehicles with unique identifiers
	for i := 0; i < p.Fleet; i++ {
		vehicles[i] = Vehicle{
			id:i,
			pos:Point{0,0},
			handledRideIDs:make([]int, 0, p.NumRides*2/p.Fleet)}
	}

	logI.Println("Done with setup!")

	// Setup done, solve problem

	for t := 0; t < p.TimeSlots; t++ {
		for _, v := range vehicles {
			if v.busy {
				if v.unavailableUntil == t {
					v.busy = false
				}
				continue
			}

			// Vehicle is idling
			for rIndex, r := range rides {
				isReachable := r.finTime -
					r.startPoint.Dist(r.endPoint) -
					v.pos.Dist(r.startPoint) > t

				if isReachable {
					v.handleRide(r, t)
					rides = append(rides[:rIndex], rides[rIndex+1:]...)
					break
					// Delete this ride, we don't need in anymore
				}

			}
		}
	}

	// Problem solved, print it
	var out bytes.Buffer
	for _, v := range vehicles {
		v.print(&out)
	}

	logI.Print("Printing data:\n" + out.String())

	// Write output to file
	err = ioutil.WriteFile("a.out", out.Bytes(), 0733)
	if err != nil {
		logE.Println(err)
	}

}

func toRide(text string, id int) Ride {
	s := strings.Split(text, " ")
	return Ride{
		id,
		Point{s2i(s[0]), s2i(s[1])},
		Point{s2i(s[2]), s2i(s[3])},
		s2i(s[4]),
		s2i(s[5]),
	}
}



func s2i (s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logE.Println(err)
	}
	return i
}

func toProblem(text string) Problem {
	s := strings.Split(text, " ")
	return Problem{
		s2i(s[0]),
		s2i(s[1]),
		s2i(s[2]),
		s2i(s[3]),
		s2i(s[4]),
		s2i(s[5]),
	}
}