package main

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"bytes"
	"io/ioutil"
	"sync"
)

var (
	logE = log.New(os.Stderr, "[ERRO] ", log.Ldate+log.Ltime+log.Ltime+log.Lshortfile)
	logI = log.New(os.Stderr, "[INFO] ", log.Ldate+log.Ltime)
	wg   sync.WaitGroup
)

type Problem struct {
	Rows      int
	Cols      int
	Fleet     int
	NumRides  int
	Bonus     int
	TimeSlots int
}

func main() {
	wg.Add(1)
	//go run("a_example.in")
	go run("b_should_be_easy.in")
	//go run("c_no_hurry.in")
	//go run("d_metropolis.in")
	//go run("e_high_bonus.in")
	wg.Wait()
}

func run(problem string) {
	defer wg.Done()
	// Start 5 threads, for every example

	file, err := os.Open("example/" + problem)
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
	}
	// Create all vehicles with unique identifiers
	for i := 0; i < p.Fleet; i++ {
		vehicles[i] = Vehicle{
			id:             i,
			pos:            Point{0, 0},
			handledRideIDs: make([]int, 0, p.NumRides*2/p.Fleet)}
	}

	logI.Println("Done with setup " + problem[0:1])
	// Setup done, solve problem

	filterCount := 0

TIME:	for time := 0; time < p.TimeSlots; time++ {
		for vIdx := range vehicles {
			vehicle := &vehicles[vIdx]

			if vehicle.unavailableUntil == time {
				vehicle.busy = false
			}
			if vehicle.busy {
				continue
			}

			filterCount+=1

			// Vehicle is idling, choose a Ride to handle
			reachableRides := Filter(rides, func(ride Ride) bool {
				return !ride.handled &&

					ride.finTime -
						ride.startPoint.Dist(ride.endPoint) -
						vehicle.pos.Dist(ride.startPoint) >= time
			})

			// Sort rides by distance from current position
			closest := func(r1, r2 *Ride) bool {
				return vehicle.pos.Dist(r1.startPoint) < vehicle.pos.Dist(r2.startPoint)
			}

			By(closest).Sort(reachableRides)
			if len(reachableRides) <	 1 {
				break TIME // Out of rides
			}
			rideToHandle := reachableRides[0] // Pick closest ride
			vehicle.handleRide(rideToHandle, time)
			rides[rideToHandle.id].handled = true // Beetje beun, maarja.
		}
		//logI.Printf("%s: Tick Tock %d", problem[0:1], time)
	}

	// Problem solved, print it
	var out bytes.Buffer
	var sumRides int = 0
	for _, v := range vehicles {
		v.print(&out)
		sumRides += len(v.handledRideIDs)
		if v.busy {
			sumRides -= 1
		}
	}

	// Stats
	logI.Printf("DONE %s", problem[0:1])

	logI.Printf("We did %d filters", filterCount)
	logI.Printf("We handled %d out of %d rides", sumRides, p.NumRides)

	// Write output to file
	err = ioutil.WriteFile("out/"+problem[:len(problem)-2]+"out", out.Bytes(), 0733)
	if err != nil {
		logE.Println(err)
	}
}

func Filter(rides []Ride, f func(Ride) bool) (fRides []Ride) {
	for _, r := range rides {
		if f(r) {
			fRides = append(fRides, r)
		}
	}
	return
}

// toRide is a helper function that parses one line from the input file to a Ride
func toRide(text string, id int) Ride {
	s := strings.Split(text, " ")
	return Ride{
		id,
		Point{stringToInt(s[0]), stringToInt(s[1])},
		Point{stringToInt(s[2]), stringToInt(s[3])},
		stringToInt(s[4]),
		stringToInt(s[5]),
		false,
	}
}

// toProblem converts the first line of the input file to the problem statement
func toProblem(text string) Problem {
	s := strings.Split(text, " ")
	return Problem{
		stringToInt(s[0]),
		stringToInt(s[1]),
		stringToInt(s[2]),
		stringToInt(s[3]),
		stringToInt(s[4]),
		stringToInt(s[5]),
	}
}

// Helper function to avoid clutter
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logE.Println(err)
	}
	return i
}
