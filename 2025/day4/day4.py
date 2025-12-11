def isRoll2(matrix, i, j):
    char = matrix[i][j]

    return 1 if char == "@" else 0


def getAllRolls():
    pass


def main():

    files = ['test.txt', 'day4.txt']

    locations = {
        "up": (0, -1),
        "down": (0, 1),
        "right": (1, 0),
        "left": (-1, 0),
        "rightup": (1, -1),
        "rightdown": (1, 1),
        "leftup": (-1, -1),
        "leftdown": (-1, 1),
    }

    for f in files:
        file = open(f'day4/{f}', "r")
        lines = file.readlines()
        global_count = 0

        for r in range(len(lines)):
            lines[r] = lines[r].strip()
            lines[r] = list(lines[r])

        while True:
            found = set()
            count = 0
            M = len(lines)
            for i in range(M):
                line = lines[i]
                N = len(line)
                for j in range(N):
                    char = line[j]
                    if char != "@":
                        continue

                    cur_count = 0
                    for d_r, d_c in locations.values():
                        n_r, n_c = i + d_r, j + d_c
                        if n_r < 0 or n_r >= len(lines) or n_c < 0 or n_c >= len(line):
                            continue
                        cur_count += isRoll2(lines, n_r, n_c)

                    if cur_count < 4:
                        found.add((i, j))
                        count += 1

            for i in range(M):
                N = len(lines[i])
                for j in range(N):
                    if (i, j) in found:
                        lines[i][j] += "x"

            if len(found) == 0:
                break
            else:
                global_count += count

        print("Part 2:", global_count)


if __name__ == "__main__":
    main()
