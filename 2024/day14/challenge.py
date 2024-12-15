import math
from dataclasses import dataclass

def resolve_star1(file) -> int:
    robots = read_file(file)

    for second in range(100):
        for robot in robots:
            robot.move()

    matrix = get_grid(robots)
    return calc_quadrant(matrix)

def resolve_star2(file) -> int:
    read_file(file)

    result = 0
    return result


def read_file (file):
    robots = []
    with open(file) as data_file:
        for line in data_file:
            parts = line.split()

            # Extract coordinates (p) and velocities (v)
            p_values = parts[0].split('=')[1]  # Extract "0,4"
            v_values = parts[1].split('=')[1]  # Extract "3,-3"

            # Convert them to integers
            coord_col, coord_row = map(int, p_values.split(','))
            vel_col, vel_row = map(int, v_values.split(','))

            # Create and return the Robot instance
            robots.append(Robot(coord_col=coord_col, coord_row=coord_row, vel_col=vel_col, vel_row=vel_row))

    return robots

max_rows = 0
max_cols = 0

@dataclass
class Robot:
    coord_col: int
    coord_row: int
    vel_col: int
    vel_row: int

    def move(self):
        if 0 <= self.coord_col + self.vel_col < max_cols:
            self.coord_col += self.vel_col
        else:
            if self.coord_col + self.vel_col < 0:
                self.coord_col = self.coord_col + max_cols + self.vel_col
            else:
                self.coord_col = self.coord_col + (self.vel_col  - max_cols)

        if 0 <= self.coord_row + self.vel_row < max_rows:
            self.coord_row += self.vel_row
        else:
            if self.coord_row + self.vel_row < 0:
                self.coord_row = self.coord_row + max_rows + self.vel_row
            else:
                self.coord_row = self.coord_row + (self.vel_row  - max_rows)


def get_grid(robots: [Robot]) -> []:
    print()
    grid = [['.' for _ in range(max_cols)] for _ in range(max_rows)]

    for robot in robots:
        value = grid[robot.coord_row][robot.coord_col]
        if value == '.':
            grid[robot.coord_row][robot.coord_col] = 1
        else:
            grid[robot.coord_row][robot.coord_col] =  value + 1


    for grid_line in grid:
        print(grid_line)

    return grid


def calc_quadrant(grid: []) -> int:
    middle_rows = math.floor(max_rows/2)
    middle_cols = math.floor(max_cols/2)

    q1 = 0
    q2 = 0
    q3 = 0
    q4 = 0

    for idx, row in enumerate(grid):
        for idy, val in enumerate(row):
            if idx == middle_rows or idy == middle_cols:
                continue

            if idx < middle_rows:
                if idy < middle_cols:
                    if val != '.':
                        q1 += val
                else:
                    if idy > middle_cols:
                        if val != '.':
                            q2 += val
            else:
                if idy < middle_cols:
                    if val != '.':
                        q3 += val
                else:
                    if idy > middle_cols:
                        if val != '.':
                            q4 += val


    print(q1, q2, q3, q4)
    return q1*q2*q3*q4
