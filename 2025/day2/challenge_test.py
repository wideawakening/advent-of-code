import unittest
import challenge


class MyTestCase(unittest.TestCase):

    def test_check_invalid_num_star1(self):
        self.assertEqual(False, challenge.check_invalid_id_star1("1"))
        self.assertEqual(False, challenge.check_invalid_id_star1("111"))
        self.assertEqual(False, challenge.check_invalid_id_star1("0101"))
        self.assertEqual(False, challenge.check_invalid_id_star1("95"))

        self.assertEqual(True, challenge.check_invalid_id_star1("11"))
        self.assertEqual(True, challenge.check_invalid_id_star1("22"))
        self.assertEqual(True, challenge.check_invalid_id_star1("99"))
        self.assertEqual(True, challenge.check_invalid_id_star1("446446"))
        self.assertEqual(True, challenge.check_invalid_id_star1("38593859"))

    def test_check_invalid_num_star2(self):
        self.assertEqual(False, challenge.check_invalid_id_star2("1"))
        self.assertEqual(False, challenge.check_invalid_id_star2("23"))
        self.assertEqual(False, challenge.check_invalid_id_star2("232023"))


        # 1 digit rep
        self.assertEqual(True, challenge.check_invalid_id_star2("11"))
        self.assertEqual(True, challenge.check_invalid_id_star2("111"))
        self.assertEqual(True, challenge.check_invalid_id_star2("222222"))
        # 2 digit rep
        # self.assertEqual(True, challenge.check_invalid_id_star2("1010"))
        self.assertEqual(True, challenge.check_invalid_id_star2("565656"))
        self.assertEqual(True, challenge.check_invalid_id_star2("1212121212"))
        self.assertEqual(True, challenge.check_invalid_id_star2("2121212121"))
        # 3 digit rep
        self.assertEqual(True, challenge.check_invalid_id_star2("446446"))
        self.assertEqual(True, challenge.check_invalid_id_star2("824824824"))
        # 4 digit rep
        self.assertEqual(True, challenge.check_invalid_id_star2("38593859"))
        # 5 digit rep
        self.assertEqual(True, challenge.check_invalid_id_star2("1188511885"))

    def test_foo(self):
        self.assertEqual(True, all(i == "aaaaaa"[0] for i in "aaaaaa"))
        self.assertEqual(False, all(i == "aaaaaa"[0] for i in "aaaaab"))

    def test_star_1_sample(self):
        result: int = challenge.resolve_star1('star_sample.txt')
        self.assertEqual(1227775554, result)

    def test_star_1_input(self):
        result: int = challenge.resolve_star1('star_input.txt')
        self.assertEqual(0, result)

    def test_star_2_sample(self):
        result: int = challenge.resolve_star2('star_sample.txt')
        self.assertEqual(4174379265, result)

    def test_star_2_input(self):
     result: int = challenge.resolve_star2('star_input.txt')
     self.assertEqual(0, result)

if __name__ == '__main__':
    unittest.main()