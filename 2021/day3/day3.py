text_file = open("day3.txt", "r")
lines = text_file.read().split("\n")


ZERO = "0"
ONE = "1"

NUM_BITS = len(lines[0])
gamma = "0b"
epsilon = "0b"

for i in range(NUM_BITS):
    zero_count = 0
    one_count = 0
    for line in lines:
        if line[i] == ZERO:
            zero_count+=1
        else:
            one_count+=1
    
    if zero_count > one_count:
        gamma+=ZERO
        epsilon+=ONE
    else: 
        gamma+=ONE
        epsilon+=ZERO

res = int(gamma, 2) * int(epsilon, 2) # base 2 int conversion



print(res)



