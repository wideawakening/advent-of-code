import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_sample(self):
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(result, 18)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(result, 0)

    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample.txt')
        self.assertEqual(result, 9)

    def test_star_2_input(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertNotEqual(result, 2401)  # too high
     self.assertNotEqual(result, 1963)  # too high
     self.assertNotEqual(result, 1927) # too high
     self.assertNotEqual(result, 1914)
     self.assertNotEqual(result, 1879)
     self.assertNotEqual(result, 1411)


# try -1 1878
     self.assertEqual(result, 0)

if __name__ == '__main__':
    unittest.main()