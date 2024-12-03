import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_sample(self):

        self.assertEqual(25, challenge.calculate("mul(5,5)"))
        self.assertEqual(25, challenge.calculate("xasdfasfmul(5,5)x"))
        self.assertEqual(125, challenge.calculate("xasdfasfmul(5,5)x1291281871mmul(10,10)"))

        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(result, 161)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(result, 0)

    # def test_star_2_sample(self):
    #     result: int = challenge.resolve_star2('star_sample.txt')
    #     self.assertEqual(result, 0)
    #
    # def test_star_2_input(self):
    #  result: int = challenge.resolve_star2('star_input.txt')
    #  self.assertEqual(result, 0)

if __name__ == '__main__':
    unittest.main()