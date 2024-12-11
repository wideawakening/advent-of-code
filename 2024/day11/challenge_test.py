import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_snippet(self):
        self.assertEqual([253000, 1, 7], challenge.blink([125,17]))
        self.assertEqual([253,0, 2024, 14168], challenge.blink([253000, 1, 7]))
        self.assertEqual([512072, 1, 20, 24, 28676032], challenge.blink([253,0, 2024, 14168]))
        self.assertEqual([512, 72, 2024, 2, 0, 2, 4, 2867, 6032], challenge.blink([512072, 1, 20, 24, 28676032]))
        self.assertEqual([1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32], challenge.blink([512, 72, 2024, 2, 0, 2, 4, 2867, 6032]))
        self.assertEqual([2097446912, 14168, 4048, 2, 0, 2, 4 ,40, 48, 2024, 40 ,48, 80, 96, 2, 8, 6 ,7, 6, 0, 3, 2], challenge.blink([1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32]))


    def test_star_1_sample(self):
        result: int = challenge.resolve_star('star_sample.txt')
        self.assertEqual(55312, result)  # this fails but input not, dig why

    def test_star_1_input(self):
        result: int = challenge.resolve_star('star_input.txt')
        self.assertEqual(0, result)

    def test_star_2_input(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertNotEqual(572595, result) # too low
     self.assertEqual(0, result)


if __name__ == '__main__':
    unittest.main()