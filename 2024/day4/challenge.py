def resolve_star1(file) -> int:
    matrix = read_file(file)
    return phrase_in_matrix(matrix, "XMAS")


def phrase_in_matrix(matrix, phrase):
    rows = len(matrix)
    cols = len(matrix[0])

    # Horizontal directions
    lr = [''.join(row) for row in matrix]  # Left-to-right
    rl = [row[::-1] for row in lr]  # Right-to-left

    # Vertical directions
    up_down = [''.join(matrix[row][col] for row in range(rows)) for col in range(cols)]  # Top-to-bottom
    down_up = [col[::-1] for col in up_down]  # Bottom-to-top

    # Check diagonals
    diagonals_lr = []
    diagonals_rl = []

    # Diagonals for top-left to bottom-right
    # First part: starting from first column of each row
    for start_row in range(rows):
        diag = []
        r, c = start_row, 0
        while r < rows and c < cols:
            diag.append(matrix[r][c])
            r += 1
            c += 1
        diagonals_lr.append(''.join(diag))

    # Second part: starting from the first row of each column (skipping the first column)
    for start_col in range(1, cols):
        diag = []
        r, c = 0, start_col
        while r < rows and c < cols:
            diag.append(matrix[r][c])
            r += 1
            c += 1
        diagonals_lr.append(''.join(diag))

    # Diagonals for top-right to bottom-left (reverse logic of top-left to bottom-right)
    # First part: starting from the top-right corner of each row
    for start_row in range(rows):
        diag = []
        r, c = start_row, cols - 1
        while r < rows and c >= 0:
            diag.append(matrix[r][c])
            r += 1
            c -= 1
        diagonals_rl.append(''.join(diag))

    # Second part: starting from the top row of each column (skipping the last column)
    for start_col in range(cols - 1):
        diag = []
        r, c = 0, start_col
        while r < rows and c >= 0:
            diag.append(matrix[r][c])
            r += 1
            c -= 1
        diagonals_rl.append(''.join(diag))

    downlr = [row[::-1] for row in diagonals_lr]
    downrl = [col[::-1] for col in diagonals_rl]


    # Check all directions
    all_directions = (lr + up_down +
                      rl + down_up +
                      diagonals_lr + diagonals_rl
                      + downlr + downrl
                      )

    # Check if the phrase appears in any direction
    count = 0
    for direction in all_directions:
        count += count_overlapping_matches(direction, phrase)
    return count

def count_overlapping_matches(text, phrase):
    count = 0
    start = 0
    while True:
        start = text.find(phrase, start)
        if start == -1:
            break
        count += 1
        start += 1  # Move one character forward to allow overlapping matches
    return count

def resolve_star2(file) -> int:
    matrix = read_file(file)
    return fucking_maria(matrix)

def fucking_maria(matrix):
    count = 0
    for row_num, row in enumerate(matrix):
        for col_num, value in enumerate(row):
            if value == 'A':
                # nice features of python matrixes that they
                if row_num -1 < 0 or col_num -1 < 0: #or row_num+1 >= len(matrix) or col_num+1 >= len(matrix[0]):
                    continue
                try:
                    if  (matrix[row_num-1][col_num-1] == 'M'
                            and matrix[row_num+1][col_num-1] == 'M'
                            and matrix[row_num-1][col_num+1] == 'S'
                            and matrix[row_num+1][col_num+1] == 'S'):
                          count +=1
                except:
                    pass
                try:
                    if  (matrix[row_num-1][col_num-1] == 'S'
                            and matrix[row_num+1][col_num-1] == 'S'
                            and matrix[row_num-1][col_num+1] == 'M'
                            and matrix[row_num+1][col_num+1] == 'M'):
                          count +=1
                except:
                    pass
                try:
                    if  (matrix[row_num-1][col_num-1] == 'S'
                            and matrix[row_num+1][col_num-1] == 'M'
                            and matrix[row_num-1][col_num+1] == 'S'
                            and matrix[row_num+1][col_num+1] == 'M'):
                        count += 1
                except:
                    pass

                try:
                    if (matrix[row_num - 1][col_num - 1] == 'M'
                            and matrix[row_num + 1][col_num - 1] == 'S'
                            and matrix[row_num - 1][col_num + 1] == 'M'
                            and matrix[row_num + 1][col_num + 1] == 'S'):
                          count +=1
                except:
                    pass

    return count


def read_file (file):
    with open(file, 'r') as f:
        lines = f.readlines()
    matrix = [list(line.strip()) for line in lines]
    return matrix

