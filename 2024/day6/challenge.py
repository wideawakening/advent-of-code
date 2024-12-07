
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
    read_file(file)

    result = 0
    return result



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

