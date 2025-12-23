import math


def resolve_star2(file, start) -> int:
    rotation_patterns = read_file(file)

    zero_rotations = 0
    for rotation_pattern in rotation_patterns:
        print(f'start: {start}, pattern: {rotation_pattern}')
        new_rotations = sample_rotation_counter(start, rotation_pattern)
        zero_rotations += new_rotations
        print(f'new_rotations: {new_rotations} \t\t\t\t\t\t\t total zero rotations: {zero_rotations}')

        start = sample_rotation(start, rotation_pattern)

    return zero_rotations


def resolve_star1(file, start) -> int:
    rotation_patterns = read_file(file)

    zero_rotations = 0
    for rotation_pattern in rotation_patterns:
        start = sample_rotation(start, rotation_pattern)
        if start == 0:
            zero_rotations += 1

    return zero_rotations



def sample_rotation(init_number, pattern_rotation) -> int:
    direction = pattern_rotation[0]
    distance = int(pattern_rotation[1:])

    if direction == 'L':
        return (init_number - distance)% 100

    if direction == 'R':
        return (init_number + distance) % 100
    return init_number


def sample_rotation_counter(init_number, pattern_rotation) -> int:
    direction = pattern_rotation[0]
    distance = int(pattern_rotation[1:])

    result = 0
    zeroes = 0

    if direction == 'L':
        result = init_number - distance
    if direction == 'R':
        result = init_number + distance

    if init_number !=0 and result <= 0:
        zeroes +=1
    zeroes += math.floor(abs(result) / 100)
    return zeroes


def read_file (file):
    rotations = []
    with open(file) as data_file:
        for data in data_file:
            rotations.append(data.strip())
    print('done')
    return rotations

