FORWARD = 'forward'
UP = "up"
DOWN = "down"


text_file = open("day2.txt", "r")
lines = text_file.read().split('\n')

depth = 0
h_distance = 0
aim = 0

def calculate_pos_change_part_1(DIR, value):
    global h_distance, depth
    if DIR == FORWARD:
        h_distance += value
    elif DIR == UP:
        depth = max(0, depth-value)
    elif DIR == DOWN:
        depth += value
    else: 
        raise Exception("SOME UNKNOWN DIRECTION: ", DIR)

def calculate_pos_change_part_2_aka_with_aim(DIR, value):
    global h_distance, depth, aim
    if DIR == FORWARD:
        h_distance += value
        depth += aim*value
    elif DIR == UP:
        aim -= value
    elif DIR == DOWN:
        aim += value
    else: 
        raise Exception("SOME UNKNOWN DIRECTION: ", DIR)

for line in lines:
    split_line = line.split(" ")
    # calculate_pos_change_part_1(split_line[0], int(split_line[1]))
    calculate_pos_change_part_2_aka_with_aim(split_line[0], int(split_line[1]))
    
print(depth, h_distance, depth*h_distance)
    


