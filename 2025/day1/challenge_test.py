import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_sample1_tests(self):
        self.assertEqual(challenge.sample_rotation(11, "R8"),19);
        self.assertEqual(challenge.sample_rotation(19, "L19"), 0);
        self.assertEqual(challenge.sample_rotation(0, "L1"), 99);
        self.assertEqual(challenge.sample_rotation(99, "R1"), 0);
        self.assertEqual(challenge.sample_rotation(5, "L10"), 95);
        self.assertEqual(challenge.sample_rotation(95, "R5"), 0);

    def test_star1_sample(self):
        result: int = challenge.resolve_star1('input_sample.txt', 50)
        self.assertEqual(result, 3)

    def test_star1_input(self):
        result: int = challenge.resolve_star1('input_star.txt', 50)
        self.assertEqual(result, 0)

    def test_sample2_tests(self):
        self.assertEqual(1, challenge.sample_rotation_counter(50, "L68"))
        self.assertEqual(1, challenge.sample_rotation_counter(95, "R60"))
        self.assertEqual(1, challenge.sample_rotation_counter(14, "L82"))
        self.assertEqual(1, challenge.sample_rotation_counter(99, "L99"))
        self.assertEqual(1, challenge.sample_rotation_counter(1, "R99"))
        self.assertEqual(10, challenge.sample_rotation_counter(50, "R1000"))
        self.assertEqual(10, challenge.sample_rotation_counter(50, "L1000"))
        self.assertEqual(0, challenge.sample_rotation_counter(0, "L5"))
        self.assertEqual(1, challenge.sample_rotation_counter(0, "L100"))
        self.assertEqual(0, challenge.sample_rotation_counter(0, "L99"))
        self.assertEqual(0, challenge.sample_rotation_counter(0, "R14"))
        self.assertEqual(1, challenge.sample_rotation_counter(0, "L198"))
        self.assertEqual(1, challenge.sample_rotation_counter(0, "R198"))

    def test_star2_sample(self):
        result: int = challenge.resolve_star2('input_sample.txt', 50)
        self.assertEqual(result, 6)

    def test_star2_sample(self):
        result: int = challenge.resolve_star2('input_sample2.txt', 50)
        self.assertEqual(result, 6)
    def test_star3_sample(self):
        result: int = challenge.resolve_star2('input_sample3.txt', 50)
        self.assertEqual(result, 2)
    def test_star2(self):
        result: int = challenge.resolve_star2('input_star.txt', 50)
        self.assertEqual(0, result)


if __name__ == '__main__':
    unittest.main()
