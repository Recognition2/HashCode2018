#!/usr/bin/env python3.6

import numpy as np
from Point import Point
from Ride import Ride
from Problem import Problem
from Vehicle import Vehicle
import time
import sys

"""
Main project for hashcode
"""


def read_file(f):
    all_data = np.loadtxt(f, dtype=int, delimiter = " ", skiprows = 0)
    first_row = all_data[0,:]
    all_data = all_data[1:all_data.shape[0],:]
    return first_row, all_data


def parse_ride(r, id):
    return Ride(id, Point(r[0], r[1]), Point(r[2], r[3]), r[4], r[5])


def main():
    # all_examples = [
    #     "a_example.in",
    #     "b_should_be_easy.in",
    #     "c_no_hurry.in",
    #     "d_metropolis.in",
    #     "e_high_bonus.in"
    # ]
    curTime = time.time()
    i = sys.argv[1]
    print(i)
    run_one(i)
    newTime = time.time()
    print("Took ", newTime - curTime, " seconds")
    print("Done with " + i)


def rideSort(rides):
    return sorted(rides, key=lambda a: a.rating())


def run_one(filename):
    # Setup
    rProbl, rRides = read_file("example/" + filename)
    rides = []

    p = Problem(rProbl[0], rProbl[1], rProbl[2], rProbl[3], rProbl[4], rProbl[5])
    vehicles = []

    i = 0
    while i < p.numRides:
        rides.append(parse_ride(rRides[i,:], i))
        i += 1

    rideToDelete = Ride()

    for i in range(0,p.numVehicles):
        vehicles.append(Vehicle(i))

    for t in range(0,p.timeSlots):
        rides = rideSort(rides)
        for v in vehicles:
            if v.busy:
                if v.unavailableUntil == 1:
                    v.busy = False
                    v.unavailableUntil = 0
                v.unavailableUntil -= 1
                continue
            for r in rides[:]:
                isReachable = r.finishTime - r.sp.dist(r.ep) - v.pos.dist(r.sp) > t
                if isReachable:
                    # Ride is reachable
                    v.handleRide(r, t)
                    rides.remove(r)
                    break



    # Print all vehicle actions
    s = ""
    for v in vehicles:
        s += v.printVehicle() + "\n"

    with open("out/" + filename, "w") as f:
        f.write(s)





if __name__ == "__main__":
    main()