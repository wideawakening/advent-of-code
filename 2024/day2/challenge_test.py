import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_star_1_sample(self):
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(result, 2)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(result, 0)

    def test_start_2_random(self):
        print()
        self.assertEqual(challenge.is_safe_row([1, 3, 2, 4, 5]), True)
        self.assertEqual(challenge.is_safe_row([8, 6, 4, 4, 1]), True)
        self.assertEqual(True, challenge.is_safe_row([49, 47, 44, 45, 44]))
        self.assertEqual(True, challenge.is_safe_row([24, 25, 24, 23, 21, 19, 18, 17]))

        self.assertEqual(False, challenge.is_safe_row([9, 12, 9, 11, 14, 16, 17, 20] ))
        self.assertEqual(False, challenge.is_safe_row([65, 68, 66, 67, 69, 70, 73, 72] ))
        self.assertEqual(False, challenge.is_safe_row([56, 58, 59, 58, 61, 64, 64]))
        self.assertEqual(True, challenge.is_safe_row([12, 10, 13, 16, 19, 21, 22]))


        self.assertEqual(False, challenge.is_safe_row([25, 26, 29, 29, 30, 33, 35, 39]))
        self.assertEqual(True, challenge.is_safe_row([88, 86, 83, 80, 79, 77, 78]))
        self.assertEqual(False, challenge.is_safe_row( [82, 85, 82, 85, 89]))
        self.assertEqual(challenge.is_safe_row([3, 6, 7, 9, 11, 8]), True)
        self.assertEqual(challenge.is_safe_row([43, 44, 45, 44, 46, 44]), False)
        self.assertEqual(challenge.is_safe_row([73, 75, 73, 74, 75, 75]), False)
    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample.txt')
        self.assertEqual(4, result)

    def test_star_2_input(self):
        result: int = challenge.resolve_star2('star_input.txt')
        self.assertNotEqual(result, 739)
        self.assertNotEqual(result, 761)
        self.assertNotEqual(result, 721)
        self.assertNotEqual(result, 709)
        self.assertNotEqual(result, 334)
        self.assertEqual(result, 0)

if __name__ == '__main__':
    unittest.main()