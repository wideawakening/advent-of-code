import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_sample(self):
        challenge.rows = 9
        challenge.cols = 10
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(41, result)

    def test_star_1_input(self):
        challenge.rows = 129
        challenge.cols = 130
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertNotEqual(5411, result)
        self.assertNotEqual(5067, result)
        self.assertNotEqual(985, result)

        self.assertEqual(0, result)


    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample.txt')
        self.assertEqual(6, result)

    def test_star_2_input(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertEqual(0, result)

if __name__ == '__main__':
    unittest.main()