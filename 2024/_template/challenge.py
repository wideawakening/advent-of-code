
def resolve_star1(file) -> int:
    read_file(file)

    result = 0
    return result

def resolve_star2(file) -> int:
    read_file(file)

    result = 0
    return result


def read_file (file):
    column1, column2 = [], []
    with open(file) as data_file:
        for data in data_file:
            num1, num2 = map(int, data.split())
            column1.append(num1)
            column2.append(num2)
    print('done')
    return column1, column2

