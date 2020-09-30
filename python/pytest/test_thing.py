from life import Life
import unittest

class TestLife(unittest.TestCase):
    def test_life_exists(self):
        life = Life(3)
        self.assertEqual(3, life.grid_width)