#!/usr/bin/env python3.7

from Point import Point

class Ride:
    """Ride"""
    id = 0

    sp = Point(0,0)
    ep = Point(0,0)

    startTime = 0 # Earliest start
    finishTime = 0 # Latest finish

    def __init__(self, id, sp, ep, st, ft):
        self.id = id
        self.sp = sp
        self.ep = ep
        self.st = st
        self.ft = ft
