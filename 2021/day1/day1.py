
EXAMPLE = [607, 618, 618, 617, 647, 716, 769, 792]
EXAMPLE_ANS = 5

text_file = open("day1.txt", "r")
lines = text_file.read().split('\n')
#  [0 1 2 3 4 5 6 7 8 ]
def getSummedSweepOfIncreased(numbers):
    increased_count = 0
    three_part_sum = []
    

    for idx in range(0, len(numbers)):
        line = int(numbers[idx])

        if len(three_part_sum) < 3:
            three_part_sum.append(line)
        else:
            if line > three_part_sum[0]:
                increased_count +=1 

            three_part_sum.pop(0)
            three_part_sum.append(line)
    return increased_count

print(getSummedSweepOfIncreased(lines)) # 1257
print(getSummedSweepOfIncreased(EXAMPLE)) #5

#  -------------------CORRECT----------------