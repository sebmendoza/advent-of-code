

# def part2(line: list[str]):

#     def backtrack(path: str, index_to_add: int):
#         nonlocal the_12_digit_max
#         if index_to_add >= len(line):
#             return
#         # path is the current number
#         if len(path) >= 12:
#             the_12_digit_max = max(the_12_digit_max, int(path))
#             return

#         # Try adding digit
#         backtrack(path+line[index_to_add], index_to_add + 1)
#         backtrack(path, index_to_add+1)

#     the_12_digit_max = 0
#     backtrack(line[0], 1)
#     print(the_12_digit_max)
#     return the_12_digit_max
# Epic

def part2(line: list[str]):
    line = line.strip()
    stack = []
    N = len(line)
    for i in range(N):
        char = line[i]
        while stack and int(stack[-1]) < int(char) and 12-len(stack) < N-i:
            stack.pop()
        stack.append(char)
        while len(stack) > 12:
            stack.pop()
    the_12_digit_max = "".join(stack)
    # print("The stack: ", the_12_digit_max)

    return int(the_12_digit_max)


def main():

    files = ['test.txt', "day3.txt"]
    for f in files:
        file = open(f'day3/{f}', "r")
        lines = file.readlines()
        res = 0
        for line in lines:
            line = line.strip()

            first_digit = line[0]
            max_joltage = 0
            for i in range(1, len(line)):
                second_digit = line[i]
                first_digit = str(first_digit)
                max_joltage = max(max_joltage, int(first_digit+second_digit))

                first_digit = int(first_digit)
                second_digit = int(second_digit)
                first_digit = max(second_digit, first_digit)

            res += max_joltage

        print("Part 1 answer: ", res)

        p2res = 0
        # 987654321111 + 811111111119 + 434234234278 + 888911112111
        for line in lines:
            p2res += part2(line)

        # part2(lines[3])
        print("Part 2: ", p2res)


if __name__ == "__main__":
    main()
