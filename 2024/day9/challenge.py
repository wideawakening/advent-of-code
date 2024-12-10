def resolve_star1(file) -> int:
    unfrag_data = list_convert_to_pre_defrag(read_file(file))
    defrag_data = list_defrag(unfrag_data)
    return list_get_checksum(defrag_data)

def list_convert_to_pre_defrag(data:str)->[]:
    defrag_data = []
    id = 0
    file_id = 0
    for c in data:
        if id==0 or id%2==0:
            for free_space in range(int(c)):
                defrag_data.append(file_id)
            file_id += 1
        else:
            if id%2 != 0:
                for free_space in range(int(c)):
                    defrag_data.append(".")
        id += 1

    return defrag_data

def list_defrag(defrag_data:list)->[str]:
    while '.' in defrag_data:
        idx_space = defrag_data.index(".")
        defrag_data[idx_space] = defrag_data[-1]
        defrag_data.pop()

    return defrag_data

def list_get_checksum(data:list) -> int:
    total = 0
    id = 0

    for num in data:
        total += int(num) * id
        id += 1
    return total











def resolve_star2(file) -> int:
    unfrag_data = list_convert_to_pre_defrag(read_file(file))
    defrag_data = block_defrag(unfrag_data)
    return list_get_checksum(defrag_data)


def block_defrag(defrag_data:list)->[str]:
    keep_defrag = True
    current_id = defrag_data[::-1][0]
    while keep_defrag or current_id < 0:

        # get fitting last block
        reversed_data = defrag_data[::-1]
        idx_current_id = defrag_data.index(current_id)
        size_block = get_continous_item_positions(current_id, reversed_data)
        if size_block == -1:
            keep_defrag = False

        # get max num of empty spaces
        idx_free = defrag_data.index(".")
        while True:
            size_free = get_continous_item_positions(".", defrag_data[idx_free:])
            if size_free == -1:
                break

            if size_free >= size_block:
                defrag_data[idx_free:idx_free+size_block] = defrag_data[idx_current_id:idx_current_id+size_block]
                defrag_data = defrag_data[:idx_current_id] + defrag_data[idx_current_id+size_block:]
                break

            idx_free += size_free

        current_id -= 1
        print(defrag_data)

    return defrag_data


def get_continous_item_positions(item: str, search_string: list) -> int:
    try:
        idx_space = search_string.index(item)
    except:
        return -1

    num_spaces = 1
    while True:
        try:
            if search_string[idx_space + num_spaces] == item:
                num_spaces += 1
            else:
                break
        except:
            break

    print(f'total num spaces of {item}: {num_spaces}')
    return num_spaces


def block_checksum(data:list) -> int:
    total = 0
    id = 0

    for num in data:
        total += int(num) * id
        id += 1
    return total


def read_file (file):
    with open(file) as file:
        return file.read()


# failed ideas once you realize IDs have more than one digit

# def ko_convert_to_pre_defrag(data:str)->str:
#     buffer = StringIO()
#     id = 0
#     file_id = 0
#     for c in data:
#         if id==0 or id%2==0:
#             for free_space in range(int(c)):
#                 buffer.write(str(file_id))
#             file_id += 1
#         else:
#             if id%2 != 0:
#                 for free_space in range(int(c)):
#                     buffer.write(".")
#         id += 1
#
#     return buffer.getvalue()
#
# def convert_to_pre_defrag(data:str)->str:
#     buffer = StringIO()
#     id = 0
#     file_id = 0
#     for c in data:
#         if id==0 or id%2==0:
#             for free_space in range(int(c)):
#                 buffer.write(f'({str(file_id)})')
#             file_id += 1
#         else:
#             if id%2 != 0:
#                 for free_space in range(int(c)):
#                     buffer.write(".")
#         id += 1
#
#     return buffer.getvalue()
#
#
#
# def ko_defrag(data:str) ->str:
#     defrag_data = list(data)
#
#     while not bool(re.fullmatch(r"^\d*\.*$", "".join(map(str, defrag_data)))):
#         idx_space = "".join(map(str, defrag_data)).index(".")
#         match = re.search(r"[^.](?=\.*$)", "".join(map(str, defrag_data)))
#         if match:
#             idx_number =  match.start()
#             val1 = defrag_data[idx_space]
#             val2 = defrag_data[idx_number]
#             defrag_data[idx_space] = val2
#             defrag_data[idx_number] = val1
#             # print("".join(map(str, defrag_data)))
#         else:
#             break
#
#     # optimization
#     result = "".join(map(str, defrag_data))
#     result = result.replace(".","")
#     print(result)
#     return result
#
#
# def get_checksum(data:str) -> int:
#     total = 0
#     id = 0
#
#     match = re.findall(r"\((\d*)\)", data)
#     for num in match:
#         total += int(num) * id
#         id += 1
#     return total
#
# def ko_get_checksum(data:str) -> int:
#     total = 0
#     id = 0
#     for c in data:
#         if c == ".":
#             break
#         total += int(c)*id
#         id += 1
#     return total

## find first '.' manually, no regexp, no strings
# def defrag(data:str) ->str:
#     defrag_data = list(data)
#
#     # while not bool(re.fullmatch(r"^\(\d*\)\.*$", data)):
#     #     data = "".join(map(str, defrag_data))
#     #     idx_space = data.index(".")
#     #     match = re.search(r"\((\d)\)$", data)
#         # if match:
#         #     number = match.group(1)
#         #     val1 = defrag_data[idx_space]
#         #     # defrag_data[idx_space] = val2
#         #     defrag_data[idx_number] = val1
#         #     print(data)
#         #     data = data.rstrip(".")
#         # else:
#         #     break
#
#     # optimization
#     print(data)
#     return data
