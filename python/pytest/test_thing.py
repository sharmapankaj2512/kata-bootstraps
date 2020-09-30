from life import Life
import unittest

# Any live cell with fewer than two live neighbours dies, as if by underpopulation.
# Any live cell with two or three live neighbours lives on to the next generation.
# Any live cell with more than three live neighbours dies, as if by overpopulation.
# Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.


class TestLife(unittest.TestCase):
    def test_life_exists(self):
        life = Life(3)
        self.assertEqual(3, life.grid_width)

    def test_a_grid_with_one_live_cell(self):
        life = Life(3)
        self.assertEqual(1, life.num_alive)

    def test_a_grid_with_two_live_cell(self):
        life = Life(3)
        self.assertEqual(2, life.num_alive)