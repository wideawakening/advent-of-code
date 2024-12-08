from reportbug.ui.text_ui import columns


# rules
# 1. frequency = single lowercase, upercase, digit
# 2. for each same frequency node-pairs, they create anti-nodes
# 3. anti-node are created, at each side, with the same distance apart than the nodes themselves
# 4. anti-nodes only withing map range


def resolve_star1(file) -> int:
    nodes = read_file(file)

    antinodes = []
    processed_permutations = []

    for frequency in nodes.keys():
        for node_a in nodes[frequency]:
            for node_b in nodes[frequency]:
                if node_a == node_b:
                    continue
                if (node_a,node_b) in processed_permutations:
                    continue

                row_distance = abs(node_a[0] - node_b[0])
                col_distance = abs(node_a[1] - node_b[1])

                # evaluate position between a/b
                ## same row
                if node_a[0] == node_b[0]:
                    if node_a[0] > node_b[0]:
                        antinode_a = (node_a[0], node_a[1] - col_distance)
                        antinode_b = (node_b[0], node_b[1] + col_distance)
                    else:
                        antinode_a = (node_a[0], node_a[1] + col_distance)
                        antinode_b = (node_b[0], node_b[1] - col_distance)

                ## same col
                if node_a[1] == node_b[1]:
                    # a > b
                    if node_a[0] > node_b[0]:
                        antinode_a = (node_a[0] - row_distance, node_a[1])
                        antinode_b = (node_b[0] + row_distance, node_b[1])
                    # a < b
                    else:
                        antinode_a = (node_a[0] + row_distance, node_a[1])
                        antinode_b = (node_b[0] - row_distance, node_b[1])

                ## a lower
                if node_a[0] < node_b[0]:
                    # a < b
                    if node_a[1] > node_b[1]:
                        antinode_a = (node_a[0] - row_distance, node_a[1] + col_distance)
                        antinode_b = (node_b[0] + row_distance, node_b[1] - col_distance)
                    else:
                        antinode_a = (node_a[0] - row_distance, node_a[1] - col_distance)
                        antinode_b = (node_b[0] + row_distance, node_b[1] + col_distance)

                ## a higher
                if node_a[0] > node_b[0]:
                    # a > b
                    if node_a[1] > node_b[1]:
                        antinode_a = (node_a[0] + row_distance, node_a[1] - col_distance)
                        antinode_b = (node_b[0] - row_distance, node_b[1] + col_distance)
                    else:
                        antinode_a = (node_a[0] + row_distance, node_a[1] + col_distance)
                        antinode_b = (node_b[0] - row_distance, node_b[1] - col_distance)


                if (node_a, node_b) not in processed_permutations:
                    if 1 <= antinode_a[0] <= max_rows and 1 <= antinode_a[1] <= max_cols:
                        if antinode_a not in antinodes:
                            antinodes.append(antinode_a)
                            print(f'appending for frequency {frequency} on {node_a},{node_b}, antinode_a {antinode_a}')

                    if 1 <= antinode_b[0] <= max_rows and 1 <= antinode_b[1] <= max_cols:
                        if antinode_b not in antinodes:
                            antinodes.append(antinode_b)
                            print(f'appending for frequency {frequency} on {node_a},{node_b}, antinode_b {antinode_b}')

                if (node_a, node_b) not in processed_permutations:
                    processed_permutations.append((node_a,node_b))
                    processed_permutations.append((node_b, node_a))

    print()
    print(antinodes)
    return len(antinodes)

def resolve_star2(file) -> int:
    read_file(file)

    result = 0
    return result



# starting with 1
max_rows=0
max_cols=0
def read_file (file):
    print()
    nodes = {}
    rows = 0

    with open(file) as data_file:
        for data in data_file:
            cols = 0
            rows += 1
            for character in data:
                cols += 1
                if character.isdigit() or character.islower() or character.isupper():
                    if character not in nodes:
                        nodes[character] = []
                    nodes[character].append((rows, cols))

    global max_rows
    max_rows = rows
    global max_cols
    max_cols = cols

    print(nodes)
    print(f'max_rows={max_rows}, max_cols={max_cols}')
    print(nodes.keys())
    return nodes

