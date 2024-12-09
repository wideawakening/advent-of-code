import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_start_1_snippet(self):
        self.assertEqual("00...111...2...333.44.5555.6666.777.888899",
                         challenge.convert_to_pre_defrag(challenge.read_file('star_sample.txt')))

        self.assertEqual(1928, challenge.get_checksum('0099811188827773336446555566..............'))

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