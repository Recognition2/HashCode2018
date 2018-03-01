#!/usr/bin/env python3.7

from Point import Point


class Ride:
    """Ride"""

    def __init__(self, id = 0, sp = Point(), ep = Point(), st = 0, ft = 0):
        self.id = id
        self.sp = sp
        self.ep = ep
        self.startTime = st
        self.finishTime = ft

    def rating(self):
        return self.ep.dist(self.sp)

