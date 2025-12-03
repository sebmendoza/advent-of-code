
from collections import deque


def part2(ranges):
    res = 0
    solutions = set()
    # Brute force solution :(
    for r in ranges:
        first, last = r.split('-')
        for i in range(int(first), int(last)+1):
            num = str(i)
            for j in range(1, len(num)//2+1):
                #  Basically check ever substring to see if it can completely make the overall string
                if len(num) % j == 0:
                    pattern = num[:j]
                    if num == pattern * (len(num) // j):
                        solutions.add(i)
                        break

            # ----- Tryied to use a queue, doesn't work for odd numbers
            # q = deque()
            # for char in num:
            #     if q and char == q[0]:
            #         q.popleft()
            #     else:
            #         q.append(char)
            # if len(q) == 0 or len(q) == 1:
            #     print(i)

    for i in solutions:
        res += i
    print(res)


def main():
    files = ['test.txt', 'day2.txt']
    for f in files:
        file = open(f'day2/{f}', 'r')
        print(f"\nOpening {f}...")
        file_text = file.read()
        ranges = file_text.split(",")
        res = 0
        for r in ranges:
            first, last = r.split('-')

            for i in range(int(first), int(last)+1):
                i = str(i)
                if len(i) % 2 == 0:
                    l = len(i)//2
                    is_double = i[:l] == i[l:]
                    if is_double:
                        res += int(i)

        print("Answer to Part1: ", res)

        part2(ranges)


if __name__ == "__main__":
    main()
