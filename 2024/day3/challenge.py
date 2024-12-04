import re

def resolve_star1(file) -> int:
    return calculate_star1(read_file(file))

def calculate_star1(content) -> int:
    total = 0
    pattern = re.compile(r'mul\((\d{1,3}),(\d{1,3})\)')

    for match in pattern.finditer(content):
        num1,num2 = match.groups()
        print(num1,num2)
        total += int(num1) * int(num2)

    return total


def resolve_star2(file) -> int:
    return calculate_star2(read_file(file))


def calculate_star2(content) -> int:
    total = 0
    on = True
    pattern = re.compile(r"mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don't\(\))")
    for found in re.findall(pattern, content):
        if "do()" in found[2]:
            on = True
        else:
            if "don't()" in found[3]:
                on = False
            else:
                if on:
                    num1, num2 = found[0], int(found[1])
                    print(num1, num2)
                    total += int(num1) * int(num2)
    return total


def read_file (file_path):
    with open(file_path) as data:
        return data.read()
