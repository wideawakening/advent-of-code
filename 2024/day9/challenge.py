import re
from io import StringIO

# disk map defrag
# defrag to free space, one at a time, from the leftmost free space block (until there are no remaining gaps)
def resolve_star1(file) -> int:
    unfrag_data = convert_to_pre_defrag(read_file(file))
    defrag_data = defrag(unfrag_data)
    return get_checksum(defrag_data)

def resolve_star2(file) -> int:
    read_file(file)

    result = 0
    return result

def convert_to_pre_defrag(data:str)->str:
    buffer = StringIO()
    id = 0
    file_id = 0
    for c in data:
        if id==0 or id%2==0:
            for free_space in range(int(c)):
                buffer.write(str(file_id))
            file_id += 1
        else:
            if id%2 != 0:
                for free_space in range(int(c)):
                    buffer.write(".")
        id += 1

    return buffer.getvalue()

def defrag(data:str) ->str:
    defrag_data = list(data)

    while not bool(re.fullmatch(r"^\d*\.*$", "".join(map(str, defrag_data)))):
        idx_space = "".join(map(str, defrag_data)).index(".")
        match = re.search(r"[^.](?=\.*$)", "".join(map(str, defrag_data)))
        if match:
            idx_number =  match.start()
            val1 = defrag_data[idx_space]
            val2 = defrag_data[idx_number]
            defrag_data[idx_space] = val2
            defrag_data[idx_number] = val1
            # print("".join(map(str, defrag_data)))
        else:
            break

    # optimization
    result = "".join(map(str, defrag_data))
    result = result.replace(".","")
    print(result)
    return result



def get_checksum(data:str) -> int:
    total = 0
    id = 0
    for c in data:
        if c == ".":
            break
        total += int(c)*id
        id += 1
    return total

def read_file (file):
    with open(file) as file:
        return file.read()

