
def resolve_star2(file) -> int:
    column1, column2 = read_file(file)

    print('column1: ', column1)
    print('column2: ', column2)

    total_similarity = 0
    for i in range(len(column1)):
        col1_number = column1[i]
        appearances = column2.count(col1_number)
        similarity = col1_number * appearances
        print(f' {col1_number}*{appearances}={similarity}')
        total_similarity += similarity

    return total_similarity

def resolve_star1(file) -> int:
    column1, column2 = read_file(file)

    print('column1: ', column1)
    print('column2: ', column2)
    column1.sort()
    column2.sort()
    print('column1: ', column1)
    print('column2: ', column2)

    total_distance = 0
    for i in range(len(column1)):
        distance = abs(column1[i] - column2[i])
        print(f' {column1[i]} {column2[i]} ; distance {distance}')
        total_distance = total_distance + distance

    return total_distance


def read_file (file):
    column1, column2 = [], []
    with open(file) as data_file:
        for data in data_file:
            num1, num2 = map(int, data.split())
            column1.append(num1)
            column2.append(num2)
    print('done')
    return column1, column2

