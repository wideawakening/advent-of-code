from typing import Dict, List, Tuple

def resolve_star1(file) -> int:
    pages_order, pages_updates = read_file(file)
    pages_updates = get_valid_updates(pages_order, pages_updates)
    return get_middle_page_numbers(pages_updates)

def get_valid_updates(pages_order:Tuple[Dict[int, List[int]]], pages_updates:List[List[int]]) -> List[List[int]]:
    valid_pages : List[List[int]] = []

    for page_sequence in pages_updates:
        if is_correct_sequence(page_sequence, pages_order):
            valid_pages.append(page_sequence)
    return valid_pages

def get_invalid_updates(pages_order:Tuple[Dict[int, List[int]]], pages_updates:List[List[int]]) -> List[List[int]]:
    invalid_pages : List[List[int]] = []

    for page_sequence in pages_updates:
        if not is_correct_sequence(page_sequence, pages_order):
            invalid_pages.append(page_sequence)
    return invalid_pages





def is_correct_sequence(page_sequence: List[int], pages_order:Tuple[Dict[int, List[int]]]) -> bool:
    reverse_page_sequence = list(reversed(page_sequence))
    for idx, checking_page in enumerate(reverse_page_sequence):  # LEARN-reverse
        if checking_page not in pages_order:
            continue
        else:
            for requirement_page in pages_order[checking_page]:
                if requirement_page in reverse_page_sequence[idx:]:
                    return False

    return True

def get_middle_page_numbers(pages_updates: List[List[int]]) -> int:
    middle_page_number_sum = 0
    for page in pages_updates:
        middle_page_number_sum += page[len(page) //2]
    return middle_page_number_sum

def resolve_star2(file) -> int:
    pages_order, pages_updates = read_file(file)
    unordered_updates = get_invalid_updates(pages_order, pages_updates)

    ordered_updates = order_updates(pages_order, unordered_updates)
    return get_middle_page_numbers(ordered_updates)


def order_updates (pages_order:Tuple[Dict[int, List[int]]], pages_updates:List[List[int]]) -> List[List[int]]:
    fixed_pages: List[List[int]] = []

    for page_sequence in pages_updates:
        fixed_sequence: List[int] = []
        for checking_page in page_sequence:
            if checking_page not in pages_order:
                fixed_sequence.append(checking_page)
            else:
                lowest_index = -1
                for requirement_page in pages_order[checking_page]:
                    if requirement_page in fixed_sequence:
                        current_index = fixed_sequence.index(requirement_page)
                        if lowest_index == -1 or current_index < lowest_index:
                            lowest_index = current_index

                if lowest_index != -1:
                    fixed_sequence.insert(lowest_index, checking_page)
                else:
                    fixed_sequence.append(checking_page)

        fixed_pages.append(fixed_sequence)

    return fixed_pages

def read_file(file: str) -> Tuple[Dict[int, List[int]], List[List[int]]]:
    pages_order: Dict[int, List[int]] = {}      # LEARN-types
    page_updates: List[List[int]] = []

    # LEARN different ways of opening file
    with open(file) as data_file:
        for line in data_file:
            if "|" in line:
                key, value = map(int, line.split("|"))
                if key not in pages_order:
                    pages_order[key] = []
                pages_order[key].append(value)

            elif "," in line:
                page_updates.append(list(map(int, line.split(","))))

    return pages_order, page_updates

