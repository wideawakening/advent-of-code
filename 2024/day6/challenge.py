
possible_paths = []
def resolve_star1(file) -> int:
    map = read_file(file)

    current_figure = '^'
    current_position_tuple = map[current_figure]
    current_position = [current_position_tuple[0][0], current_position_tuple[0][1]]  #row, col
    map['X'] = []
    map['X'].append([current_position[0], current_position[1]])
    current_position[0] = current_position[0] - 1
    map['X'].append([current_position[0], current_position[1]])


    print()
    while not is_exit(current_figure, current_position):

        if is_next_position_blocked(current_position, current_figure, map):
            match current_figure:
                case '^':
                    current_figure = '>'
                    print('changed direction to >')
                case 'v':
                    current_figure = '<'
                    print('changed direction to <')
                case '>':
                    current_figure = 'v'
                    print('changed direction to v')
                case '<':
                    current_figure = '^'
                    print('changed direction to ^')

        match current_figure:
            case '^':
                current_position[0] = current_position[0] - 1
            case 'v':
                current_position[0] = current_position[0] + 1
            case '>':
                current_position[1] = current_position[1] + 1
            case '<':
                current_position[1] = current_position[1] - 1

        if current_position not in map['X']:
            map['X'].append([current_position[0], current_position[1]])
            print(current_position)

    print()
    print(map['X'])
    possible_paths.extend(map['X'])
    return len(map['X'])

def is_next_position_blocked(current_position, current_figure, map) -> bool:
    next_position = ()
    match current_figure:
        case '^':
            next_position = (current_position[0] - 1, current_position[1])
        case 'v':
            next_position = (current_position[0] + 1, current_position[1])
        case '>':
            next_position = (current_position[0], current_position[1] + 1)
        case '<':
            next_position = (current_position[0], current_position[1] - 1)
    return next_position in map['#']

cols = 0
rows = 0
def is_exit(current_figure: str, current_position: []) -> bool:
    match current_figure:
        case '^':
            return current_position[0] == 0
        case 'v':
            return current_position[0] == rows
        case '>':
            return current_position[1] == cols
        case '<':
            return current_position[1] == 0
    return False




def resolve_star2(file) -> int:
    map = read_file(file)

    print()
    numloops = 0

    # first run resolve_star1() to get possible paths and then get the 4 borders
    resolve_star1(file)
    min_row, max_row, min_col, max_col = 999, 0, 999, 0
    for possible_path in possible_paths:
        if possible_path[0] < min_row:
            min_row = possible_path[0]
        if max_row < possible_path[0]:
            max_row = possible_path[0]
        if possible_path[1] < min_col:
            min_col = possible_path[1]
        if max_col < possible_path[1]:
            max_col = possible_path[1]
    print(f'{min_row}-{max_row}-{min_col}-{max_col}')

    # copy original blockers
    original_blocks =  list(map['#'])

    for row in range(min_row, max_row):
        for col in range(min_col, max_col):
            map['#'] = list(original_blocks)
            steps = 0
            if (row, col) not in map['#']:
                map['#'].append((row,col))

            current_figure = '^'
            current_position_tuple = map[current_figure]
            current_position = [current_position_tuple[0][0], current_position_tuple[0][1]]  # row, col
            current_position[0] = current_position[0] - 1

            while not is_exit(current_figure, current_position) and numloops < 2000: # too high result
                if is_next_position_blocked(current_position, current_figure, map):
                    match current_figure:
                        case '^':
                            current_figure = '>'
                        case 'v':
                            current_figure = '<'
                        case '>':
                            current_figure = 'v'
                        case '<':
                            current_figure = '^'

                match current_figure:
                    case '^':
                        current_position[0] = current_position[0] - 1
                    case 'v':
                        current_position[0] = current_position[0] + 1
                    case '>':
                        current_position[1] = current_position[1] + 1
                    case '<':
                        current_position[1] = current_position[1] - 1

                # simplify star1 code
                steps +=1

                if steps > 100_000:
                    numloops += 1
                    break


    print()
    return numloops



def read_file (file):
    print()
    map = {}
    with open(file, 'r') as file:
        for row_index, line in enumerate(file):
            for col_index, str in enumerate(line.strip()):
                if str not in map:
                    map[str] = []
                map[str].append((row_index, col_index))
    # for str, pos in map.items():
    #     print(f"'{str}': {pos}")
    return map

