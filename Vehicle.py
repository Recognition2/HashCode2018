

from Ride import Ride
from Point import Point
from enum import Enum

class Vehicle:
    id = 0


    def __init__(self, id):
        self.id = id
        self.pos = Point(0, 0)
        self.ridesHandled = []
        self.busy = False
        self.unavailableUntil = 0

    def handleRide(self, r, t):
        self.busy = True
        self.unavailableUntil = self.pos.dist(r.sp) + r.sp.dist(r.ep)
        self.ridesHandled.append(r.id)
        self.pos = r.ep

    def printVehicle(self):
        s = str(len(self.ridesHandled)) + " "
        for v in self.ridesHandled:
            s += str(v) + " "
        # if self.busy:
        #     s = " ".join(s.split(" ")[:-1])
        return s

