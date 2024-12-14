import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_snippet1(self):
        challenge.max_rows = 7
        challenge.max_cols = 8
        result: int = challenge.resolve_star1('snippet1.txt')
        self.assertEqual(2, result)

    def test_star_1_snippet2(self):
        challenge.max_rows = 7
        challenge.max_cols = 8
        result: int = challenge.resolve_star1('snippet2.txt')
        self.assertEqual(4, result)

    def test_star_1_snippet2(self):
        challenge.max_rows = 7
        challenge.max_cols = 8
        result: int = challenge.resolve_star1('snippet3.txt')
        self.assertEqual(3, result)

    def test_star_1_sample(self):
        challenge.max_rows = 8
        challenge.max_cols = 9
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(36, result)

    def test_star_1_input(self):
        challenge.max_rows = 47
        challenge.max_cols = 48
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertNotEqual(909 , result) # too high
        self.assertEqual(0, result)

    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample.txt')
        self.assertEqual(0, result)

    def test_star_2_input(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertEqual(0, result)

if __name__ == '__main__':
    unittest.main()