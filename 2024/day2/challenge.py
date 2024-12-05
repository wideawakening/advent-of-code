

# one report per line
# each report/line has levels
# safe report = values increasing or decreasing
# adjacent levels differ 1..3

def resolve_star1(file) -> int:
    report = read_file(file)
    print(report)

    result = 0
    for i, row in enumerate(report):
        previous_val = 0
        # print(f"sorted", sorted(row))
        # print(f"sorted", sorted(row, reverse=True))

        if sorted(row) != row and sorted(row,reverse=True) != row:
            continue

        for j, number in enumerate(row):
            if previous_val == 0:
                previous_val = number

            #print(f"data[{i}][{j}] = {number}")
            if abs(number - previous_val) <= 3 and (j==0 or abs(number - previous_val) != 0):
                previous_val = number
            else:
                previous_val = -1
                print(f"invalid report: {row}")
                break

        if previous_val != -1:
            print(f"valid report: {row}")
            result +=1

    return result



def resolve_star2(file) -> int:
    report = read_file(file)
    #print(report)

    result = 0
    for i, row in enumerate(report):
        if is_safe_row(row):
            result +=1
    return result

def is_safe_row(row) -> bool:
    ok = is_ok_withouth_one(row) #or is_ok_within_range(row)
    if ok:
        print(f"safe report: {row}")
    else:
        print(f"unsafe report: {row}")

    return ok

def is_ok_within_range(row) -> bool:

    dampeable = False
    order_type = 0 # 1 asc -1 desc

    for idx, number in enumerate(row):
        try:
            next_num = row[idx+1]

            # check asc|desc
            new_order = -1 if number < next_num else (+1 if number > next_num else 0)
            if order_type != 0 and new_order != order_type:
                return False
            else:
                order_type = new_order

            diff = abs(next_num - number)
            if diff == 0:
                if not dampeable:
                    return False
                else:
                    dampeable = False
                continue


            if  diff < 1 or diff > 3:
                if not dampeable:
                    print(f"unsafe report: {row}")
                    return False
                else:
                    dampeable = False
        except:
            pass


    print(f"safe report: {row}")
    return True



def is_ok_withouth_one(row) -> bool:
    # remove one number for list, and check if its ordered ok
    for i, number in enumerate(row):
        subreport = row[:i] + row[i + 1:]

        is_sorted = (
                sorted(subreport) == subreport or
                sorted(subreport, reverse=True) == subreport
        )
        if not is_sorted:
             continue

        if is_ok_within_range(subreport):
            return True

    return False

def is_ok_within_range(row) -> bool:

    dampeable = False
    order_type = 0 # 1 asc -1 desc

    for idx, number in enumerate(row):
        try:
            next_num = row[idx+1]

            # check asc|desc
            new_order = -1 if number < next_num else (+1 if number > next_num else 0)
            if order_type != 0 and new_order != order_type:
                return False
            else:
                order_type = new_order

            diff = abs(next_num - number)
            if diff == 0:
                if not dampeable:
                    return False
                else:
                    dampeable = False
                continue


            if  diff < 1 or diff > 3:
                if not dampeable:
                    print(f"unsafe report: {row}")
                    return False
                else:
                    dampeable = False
        except:
            pass


    print(f"safe report: {row}")
    return True


# def resolve_star2(file) -> int:
#     report = read_file(file)
#     #print(report)
#
#     result = 0
#     for i, row in enumerate(report):
#         if is_safe_row(row):
#             result +=1
#     return result
#
# def is_safe_row(row) -> bool:
#     print("---")
#     previous_val = 0
#     dampener = 0
#
#     if not check_is_damped_sorted(row):
#         print(f"unsafe report: {row}")
#         return False
#
#     for j, number in enumerate(row):
#         if previous_val == 0:
#             previous_val = number
#
#         # print(f"data[{i}][{j}] = {number}")
#         if abs(number - previous_val) <= 3 and (j == 0 or abs(number - previous_val) != 0):
#             previous_val = number
#         else:
#             if dampener == 0:
#                 dampener = 1
#                 #print("dampener")
#             else:
#                 print(f"unsafe report: {row}")
#                 return False
#
#     if previous_val != -1:
#         print(f"safe report: {row}")
#         return True
#
#     return False
#
#
# def check_is_damped_sorted(row):
#     for i, number in enumerate(row):
#         is_sorted = (
#                 sorted(row[:i] + row[i + 1:]) == row[:i] + row[i + 1:] or
#                 sorted(row[:i] + row[i + 1:], reverse=True) == row[:i] + row[i + 1:]
#         )
#         if is_sorted:
#             return True
#     return False
#

#
# def check_is_damped_sorted(row):
#     return check_damped_sorted_internal(row, False) or check_damped_sorted_internal(row, True)
#
# def check_damped_sorted_internal(row, reverse=False) -> bool:
#     if check_is_damped_sorted_withouth_0(row, reverse) :
#         return True
#
#     damped = False
#     previous_val = 0
#     for i, number in enumerate(row):
#         if i!=0 and (reverse== False and number >= previous_val) or (reverse == True and number <= previous_val):
#             if not damped:
#                 damped = True
#                 continue
#             else:
#                 return False
#         previous_val = number
#
#     return True
#
# def check_is_damped_sorted_withouth_0(row, reverse=False) -> bool:
#     previous_val = row[1]
#     for i, number in enumerate(row):
#         if i < 2:
#             continue
#         if (reverse== False and number >= previous_val) or (reverse == True and number <= previous_val):
#             return False
#         previous_val = number
#     return True



def read_file (file):
    reports = []
    with open(file) as data_file:
        for line in data_file:
            levels = list(map(int, line.split()))
            reports.append(levels)
    return reports

