from sqlalchemy import false


def resolve_star1(file) -> int:
    content = read_file(file)

    sum_invalid_ids = 0
    for content_item in content:
        rangesAB = content_item.split("-")
        rangeA = rangesAB[0]
        rangeB = rangesAB[1]

        # if check_invalid_id(rangeA):
        #     sum_invalid_ids += int(rangeA)
        #     print(rangeA)
        if check_invalid_id_star1(rangeB):
            sum_invalid_ids += int(rangeB)
            print(rangeB)

        for num_rangeN in range(int(rangeA), int(rangeB)):
            if check_invalid_id_star1(str(num_rangeN)):
                sum_invalid_ids += int(num_rangeN)
                print(num_rangeN)

    return sum_invalid_ids

def resolve_star2(file) -> int:
    content = read_file(file)

    sum_invalid_ids = 0
    for content_item in content:
        rangesAB = content_item.split("-")
        rangeA = rangesAB[0]
        rangeB = rangesAB[1]

        # if check_invalid_id(rangeA):
        #     sum_invalid_ids += int(rangeA)
        #     print(rangeA)
        if check_invalid_id_star2(rangeB):
            sum_invalid_ids += int(rangeB)
            print(rangeB)

        for num_rangeN in range(int(rangeA), int(rangeB)):
            if check_invalid_id_star2(str(num_rangeN)):
                sum_invalid_ids += int(num_rangeN)
                print(num_rangeN)

    return sum_invalid_ids

def check_invalid_id_star1(num) -> bool:
    if len(num)%2 !=0:
        return False

    if int(num[0]) == 0:
        return False

    numA = num[:int((len(num)/2))]
    numB = num[int(len(num) / 2):]

    return numA == numB


def check_invalid_id_star2(num) -> bool:
    if int(num[0]) == 0:
        return False

    is_invalid = False
    for digits in range (1,len(num)):   # TODO optimize and do less nums
        chunks = [num[i:i + digits] for i in range(0, len(num), digits)]
        is_invalid = is_invalid or all(chunk == chunks[0] for chunk in chunks)

    return is_invalid





def read_file (file):
    content = []
    with open(file) as data_file:
        for data in data_file:
            for item in data.split(","):
                content.append(item)
    print('done')
    return content

