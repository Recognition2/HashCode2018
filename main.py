#!/usr/bin/env python3.6

import numpy as np
from Point import Point
from Ride import Ride
from Problem import Problem
from Vehicle import Vehicle

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
    # Setup
    filename = "example/a_example.in"
    rProbl, rRides = read_file(filename)
    rides = []

    p = Problem(rProbl[0], rProbl[1], rProbl[2], rProbl[3], rProbl[4], rProbl[5])
    vehicles = []

    i = 0
    while i < p.numRides:
        rides.append(parse_ride(rRides[i,:], i))
        i += 1

    for i in range(0,p.numVehicles):
        vehicles.append(Vehicle(i))

    for t in range(0,p.timeSlots):
        for v in vehicles:
            if v.busy:
                if v.unavailableUntil == 0:
                    v.busy = False
                v.unavailableUntil -= 1
                continue
            for r in rides:
                if r.startTime + v.Pos.dist(r) > t:
                    # Ride is reachable
                    v.



    # Print all vehicle actions
    s = ""
    for v in vehicles:
        s += v.print() + "\n"

    with open("output.txt", "w") as f:
        f.write(s)





if __name__ == "__main__":
    main()