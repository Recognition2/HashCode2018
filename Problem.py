
from Ride import Ride
from Point import Point


class Problem:
    rows = 0
    cols = 0
    numVehicles = 0
    numRides = 0
    bonus = 0
    timeSlots = 0

    def __init__(self, rows, cols, fleet, numRides, bonus, timeSlots):
        self.rows = rows
        self.cols = cols
        self.numVehicles = fleet
        self.numRides = numRides
        self.bonus = bonus
        self.timeSlots = timeSlots
