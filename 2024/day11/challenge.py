from functools import cache

def resolve_star(file) -> int:
    stones = read_file(file)
    print()
    print(stones)

    for i in range(25):
        stones = blink(stones)

    return len(stones)

def blink(stones:list[int]) -> list[int]:
    new_stones = []
    for idx, num in enumerate(stones):
        if num == 0:
            new_stones.append(1)
            continue

        if len(str(num))%2 == 0:
            txt_num = str(num)
            new_stones.append(int(txt_num[:int(len(txt_num) / 2)]))
            new_stones.append(int(txt_num[int(len(txt_num) / 2):]))
            continue

        new_stones.append(num*2024)

    # print(new_stones)
    return new_stones

def resolve_star2(file) -> int:
    stones = read_file(file)
    print()
    print(stones)

    total = 0
    for stone in stones:
        total += count_stones(stone, 75)

    return total



@cache
def count_stones(num:int, blink_num:int) -> int:

    if blink_num == 0:
        return 1

    if num == 0:
        return count_stones(1, blink_num-1)

    if len(str(num))%2 == 0:
        txt_num = str(num)
        return (count_stones(int(txt_num[:int(len(txt_num) / 2)]),blink_num-1)
                + count_stones(int(txt_num[int(len(txt_num) / 2):]),blink_num-1))

    return count_stones(num*2024,blink_num-1)



def read_file (file)-> list:
    with open(file) as data_file :
        stones = []
        for data in data_file:
            for stone in data.split():
                stones.append(int(stone))
    return stones

