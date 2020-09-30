import numpy as np


class Life():

    def __init__(self, grid_width):
        self.grid_width = grid_width
        self.num_alive = 1
        self.grid = np.array(grid_width, grid_width)


