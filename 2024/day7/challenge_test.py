import unittest
import challenge
from day7.challenge import has_possible_operation


class MyTestCase(unittest.TestCase):

    def test_star_1_sample(self):
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(3749, result)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(0, result)

    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample.txt')
        self.assertEqual(11387, result)

    def test_star_2_samples(self):
        self.assertEqual(False, has_possible_operation(100, [50,50,2]))

    def test_star_2_input(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertNotEqual(264184082159178, result) # too high #concat_numbers # int(str(),str())
     self.assertNotEqual(3122337891434, result) #  current_result*10^math.ceil(math.log(operation_value, 10))+operation_value
     self.assertNotEqual(3119477785044, result)
     self.assertNotEqual(3122183584666, result)  # too low

     self.assertEqual(0, result)

if __name__ == '__main__':
    unittest.main()