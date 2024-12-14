
# 0..9; +1 increment
# up/down/left/right

max_rows = 0
max_cols = 0

def resolve_star1(file) -> int:
    print()
    map = read_file(file)
    cords_start = [(row_idx, col_idx) for row_idx, row in enumerate(map) for col_idx, value in enumerate(row) if value == 0]

    global tops
    total_score = 0
    for cord_current in cords_start:
        print('start point is', cord_current)
        score = do_trailhead(0, cord_current, map, [])
        print('score is ', len(tops))
        total_score += len(tops)
        tops = []
    return total_score


tops = []
def do_trailhead(level: int, cord_current: tuple, map: list, path:list) -> int:
    print('level is', level, 'cord_current', cord_current, 'path: ', path)
    score = 0
    if level == 9: #and is_top(cord_current, map):
        if (cord_current[0], cord_current[1]) not in tops:
            tops.append((cord_current[0], cord_current[1]))
        return 1

    # check cord_current options
    # check within map and sequential value

    ## up
    next_step = (cord_current[0]-1, cord_current[1])
    if is_within_map(cord_current) and is_trailable(level, next_step, map):
        path.append(next_step)
        score += do_trailhead(level + 1, next_step, map, list(path))

    ## down
    next_step = (cord_current[0]+1, cord_current[1])
    if is_within_map(cord_current) and is_trailable(level, next_step, map):
        path.append(next_step)
        score += do_trailhead(level + 1, next_step, map, list(path))

    ## left
    next_step = (cord_current[0], cord_current[1]-1)
    if is_within_map(cord_current) and is_trailable(level, next_step, map):
        path.append(next_step)
        score += do_trailhead(level + 1, next_step, map, list(path))

    ## right
    next_step = (cord_current[0], cord_current[1]+1)
    if is_within_map(cord_current) and is_trailable(level, next_step, map):
        path.append(next_step)
        score += do_trailhead(level + 1, next_step, map, list(path))

    return score    # score is not used


def is_within_map(position:tuple) -> bool:
    if (0 <= position[0] < max_rows
            and 0 <= position[1] < max_cols):
        return True
    return False

def is_trailable(score:int, next_step:tuple, map:list) -> bool:
    try:
        if score +1 == map[next_step[0]][next_step[1]]:
            return True
    except:
        pass
    return False






def resolve_star2(file) -> int:
    read_file(file)
    result = 0
    return result



def read_file (file) -> list:
    # with open(file) as file:
    #     content = file.read()

    # data = [list(map(any, line)) for line in content.strip().splitlines()]
    # return data

    with open(file, 'r') as content:
        lines = content.readlines()
    matrix = []
    for line in lines:
        row = []
        for char in line.strip():
            if char == '.':
                row.append(-1)
            else:
                row.append(int(char))
        matrix.append(row)

    # Print the resulting matrix
    for row in matrix:
        print(row)
    return matrix