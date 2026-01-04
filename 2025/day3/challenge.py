
def resolve_star1(file) -> int:
    data = read_file(file)

    result = 0
    for line in data:
        result += get_biggest_nums(line)
    return result

def get_biggest_nums(data):

    # first biggest
    # we don't use the last character since it may not have a consecutive. ex.: 811111111111119
    num_biggest = 0
    pos_biggest = 0
    for i in range(0,len(data)-1):
        num_i = int(data[i])
        if num_i > num_biggest:
            pos_biggest = i
            num_biggest = num_i
            print(f"found biggest {data[i]} in {i}")


    # second consecutive biggest
    next_biggest = 0
    num_next_biggest = 0
    for i in range(pos_biggest+1,len(data)):
        num_i = int(data[i])
        if num_i > num_next_biggest:
            next_biggest = i
            num_next_biggest = num_i
            print(f"found biggest {data[i]} in {i}")

    return int(f"{num_biggest}{num_next_biggest}")

def resolve_star1(file) -> int:
    data = read_file(file)

    result = 0
    for line in data:
        result += get_biggest_nums2(line)
    return result

def get_biggest_nums2(data):
    data_biggest = [0] * len(data)

    num_biggest = 0
    pos_biggest = 0
    for j in range(0,12):
        # get biggest num
        for i in range(pos_biggest, len(data)):

            num_i = int(data[i])
            if data_biggest[i] == 0 and num_i > num_biggest:
                pos_biggest = i
                num_biggest = num_i

            print(f"found biggest {num_biggest} in {pos_biggest}")
            data_biggest[pos_biggest] = num_biggest

            # reset if we reached the end
            if i + 2 >= len(data):
                num_biggest = 0
                pos_biggest = 0

    return int(''.join(str(d) for d in list(filter(lambda x: x != 0, data_biggest))))

def read_file (file):
    content = []
    with open(file) as data_file:
        for data in data_file:
            content.append(data.strip())
    return content

