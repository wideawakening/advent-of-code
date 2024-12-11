
# 0..9; +1 increment
# up/down/left/right

max_rows = 0
max_cols = 0

def resolve_star1(file) -> int:
    print()
    map = read_file(file)
    cords_start = [(row_idx, col_idx) for row_idx, row in enumerate(map) for col_idx, value in enumerate(row) if value == 0]

    total_score = 0
    for cord_current in cords_start:
        print('start point is', cord_current)
        score = do_trailhead(0, cord_current, map, [])
        print('score is ', score)
        total_score += score
    return total_score


def do_trailhead(level: int, cord_current: tuple, map: list, path:list) -> int:
    print('level is', level, 'cord_current', cord_current)
    score = 0
    if level == 8 and is_top(cord_current, map):
        return 1

    # check cord_current options
    # check within map and sequential value

    ## up
    next = (cord_current[0]-1, cord_current[1])
    path.append(next)
    if is_within_map(cord_current) and is_trailable(level, next, map):
        score += do_trailhead(level + 1, next, map, path)

    ## down
    next = (cord_current[0]+1, cord_current[1])
    path.append(next)
    if is_within_map(cord_current) and is_trailable(level, next, map):
        score += do_trailhead(level + 1, next, map, path)

    ## left
    next = (cord_current[0], cord_current[1]-1)
    path.append(next)
    if is_within_map(cord_current) and is_trailable(level, next, map):
        score += do_trailhead(level + 1, next, map, path)

    ## right
    next = (cord_current[0], cord_current[1]+1)
    path.append(next)
    if is_within_map(cord_current) and is_trailable(level, next, map):
        score += do_trailhead(level + 1, next, map, path)

    return score


def is_within_map(position:tuple) -> bool:
    if 0 <= position[0] < max_rows:
        return True
    if 0 <= position[1] < max_cols:
        return True
    return False

def is_trailable(score:int, position:tuple, map:list) -> bool:
    try:
        if score +1 == map[position[0]][position[1]]:
            return True
    except:
        pass
    return False

def is_top(cord_current, map)->bool:
    top = False
    try:
        if  map[cord_current[0]-1][cord_current[1]] == 9:
            top = True
    except:
        pass
    try:
        if map[cord_current[0] + 1][cord_current[1]] == 9:
            top = True
    except:
        pass
    try:
        if map[cord_current[0]][cord_current[1]-1] == 9:
            top = True
    except:
        pass
    try:
        if map[cord_current[0]][cord_current[1]+1] == 9:
           top = True
    except:
        pass

    return top


def resolve_star2(file) -> int:
    read_file(file)
    result = 0
    return result



def read_file (file) -> list:
    with open(file) as file:
        content = file.read()

    data = [list(map(int, line)) for line in content.strip().splitlines()]
    return data

