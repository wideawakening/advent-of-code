def resolve_star1(file) -> int:
    operations = read_file(file)

    result = 0
    for operation_result in operations:
        if has_possible_operation(operation_result, operations[operation_result]):
            result += operation_result
    return result

def has_possible_operation(operation_result:int, operation_values:list[int]) -> bool:

    def operate(current_result: int, operation_result:int, operation_values:list[int]) -> bool:
        if current_result == operation_result:
            return True

        if len(operation_values) == 0:
            return False

        if len(operation_values) == 1 and operation_result == operation_values[0]:
            return True

        return (
            operate(current_result + operation_values[0], operation_result, operation_values[1:]) or
            operate(current_result * operation_values[0], operation_result, operation_values[1:])
            )

    return operate(0, operation_result, operation_values)

def resolve_star2(file) -> int:
    read_file(file)

    result = 0
    return result



def read_file (file):
    print()
    operations = {}
    with open(file) as data_file:
        for data in data_file:
            result_txt, numbers  = map(str, data.split(":"))
            result = int(result_txt)
            if result not in operations:
                operations[result] = []
            for number in numbers.split():
                operations[result].append(int(number))

    print(operations)
    return operations

