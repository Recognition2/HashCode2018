

from Ride import Ride
from Point import Point
from enum import Enum

class Vehicle:
    id = 0
    pos = Point(0, 0)
    ridesHandled = []
    busy = False
    unavailableUntil = 0

    def __init__(self, id):
        self.id = id

    def print(self):
        s = ""
        for v in self.ridesHandled:
            s += v + " "
        return s