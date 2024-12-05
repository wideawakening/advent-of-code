import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_sample(self):

        self.assertEqual(25, challenge.calculate_star1("mul(5,5)"))
        self.assertEqual(25, challenge.calculate_star1("xasdfasfmul(5,5)x"))
        self.assertEqual(125, challenge.calculate_star1("xasdfasfmul(5,5)x1291281871mmul(10,10)"))

        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(161, result)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(0, result)

    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample_2.txt')
        self.assertEqual(48, result)


    def test_star_2_input(self):
        result: int = challenge.resolve_star2('star_input.txt')
        self.assertNotEqual(3032951, result)
        self.assertNotEqual(5018915, result)
        self.assertEqual(0, result)

if __name__ == '__main__':
    unittest.main()