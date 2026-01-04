import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_get_biggest_nums(self):
        self.assertEqual(98, challenge.get_biggest_nums("987654321111111"))
        self.assertEqual(89, challenge.get_biggest_nums("811111111111119"))
        self.assertEqual(78, challenge.get_biggest_nums("234234234234278"))
        self.assertEqual(92, challenge.get_biggest_nums("818181911112111"))


    def test_star_1_sample(self):
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(357, result)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(0, result)

    def test_get_biggest_nums2(self):
        self.assertEqual(987654321111, challenge.get_biggest_nums2("987654321111111"))
        # self.assertEqual(811111111119, challenge.get_biggest_nums2("811111111111119"))
        # self.assertEqual(434234234278, challenge.get_biggest_nums2("234234234234278"))
        # self.assertEqual(888911112111, challenge.get_biggest_nums2("818181911112111"))

    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample.txt')
        self.assertEqual(3121910778619, result)

    def test_star_2_input(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertEqual(0, result)

if __name__ == '__main__':
    unittest.main()