import re

def resolve_star1(file) -> int:
    return calculate(read_file(file))

# todo? trim
def calculate(content) -> int:
    total = 0
    pattern = re.compile(r'mul\((\d{1,3}),(\d{1,3})\)')  # Groups for area, prefix, and line number

    for match in pattern.finditer(content):
        num1,num2 = match.groups()
        print(num1,num2)
        total += int(num1) * int(num2)

    return total


def resolve_star2(file) -> int:
    result = 0
    return result


def read_file (file_path):
    with open(file_path) as data:
        return data.read()
