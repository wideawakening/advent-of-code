import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_snippet1(self):
        challenge.max_rows = 7
        challenge.max_cols = 11
        result: int = challenge.resolve_star1('star1_snippet1.txt')


    def test_star_1_sample(self):
        challenge.max_rows = 7
        challenge.max_cols = 11
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(12, result)

    def test_star_1_input(self):
        challenge.max_rows = 103
        challenge.max_cols = 101
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertNotEqual(216859104, result) # too low
        self.assertEqual(0, result)


    def test_star_2_input(self):
        challenge.max_rows = 103
        challenge.max_cols = 101
        result: int = challenge.resolve_star2('star_input.txt')
        self.assertEqual(0, result)

if __name__ == '__main__':
    unittest.main()