import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_start_1_snippet(self):
        self.assertEqual([0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6], challenge.list_defrag([0, 0, '.', '.', '.', 1, 1, 1, '.', '.', '.', 2, '.', '.', '.', 3, 3, 3, '.', 4, 4, '.', 5, 5, 5, 5, '.', 6, 6, 6, 6, '.', 7, 7, 7, '.', 8, 8, 8, 8, 9, 9]))


        self.assertEqual("00...111...2...333.44.5555.6666.777.888899",
                         challenge.ko_convert_to_pre_defrag(challenge.read_file('star_sample.txt')))
        self.assertEqual("(0)(0)...(1)(1)(1)...(2)...(3)(3)(3).(4)(4).(5)(5)(5)(5).(6)(6)(6)(6).(7)(7)(7).(8)(8)(8)(8)(9)(9)",
                         challenge.convert_to_pre_defrag(challenge.read_file('star_sample.txt')))

        self.assertEqual(1928, challenge.ko_get_checksum('0099811188827773336446555566..............'))
        self.assertEqual(1928, challenge.list_get_checksum([0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6]))


    def test_star_1_sample(self):
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(1928, result)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(0, result)

    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample.txt')
        self.assertEqual(0, result)

    def test_star_2_input(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertEqual(0, result)

if __name__ == '__main__':
    unittest.main()