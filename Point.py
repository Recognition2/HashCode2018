
class Point:
    row = 0
    col = 0

    def __init__(self, row = 0, col = 0):
        self.row = row
        self.col = col

    def dist(self, p):
        return abs(self.col - p.col) + abs(self.row - p.row)
