import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_sample(self):
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(result, 11)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(result, 0)

    def test_star_2_sample(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertEqual(result, 0)

if __name__ == '__main__':
    unittest.main()
