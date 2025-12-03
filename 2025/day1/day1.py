"""
Complete:
Day   -Part 1-   -Part 2-
  1   22:55:10       >24h
"""


# Solution attempt that did not work
# for i in lines:
#     direction = i[0]
#     steps = int(i[1:])
#     fromzero = position == 0

#     if direction == "R":
#         position += steps
#         if position > mod and position-steps != mod-1:
#             password += 1
#             if steps // mod > 1:
#                 password += steps//mod
#         position = position % mod
#         if position == 0:
#             password += 1

#     elif direction == "L":
#         position -= steps
#         if position < 0 and not fromzero:
#             password += 1
#             if steps // mod > 1:
#                 password += steps//mod

#         if position < 0:
#             position = (mod-(position*-1 % mod)) % mod
#         else:
#             position = position % mod

#         if position == 0:
#             password += 1

#     else:
#         return ValueError("Direction is not L or R")


"""
Date: Dec 1, 2025
Notes:
    The actual password is the number of times the dial is
    left pointing at 0 after any rotation in the sequence.
"""


init_position = 50
mod = 100


def main():

    files = ['test.txt', 'test1.txt', 'test1.txt', 'day1.txt']
    for f in files:
        file = open(f'day1/{f}', 'r')
        print(f"\nOpening {f}...")
        lines = file.readlines()
        position = init_position
        password = 0

        for i in lines:
            direction = i[0]
            steps = int(i[1:])
            while steps > 0:
                if direction == "L":
                    position -= 1
                else:
                    position += 1

                if position == 100:
                    position = 0
                if position == -1:
                    position = 99
                if position == 0:
                    password += 1
                # print(position)
                steps -= 1

        file.close()
        print(password)


if __name__ == '__main__':
    main()
