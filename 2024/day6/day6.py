from typing import List, Set, Tuple, Dict
from enum import Enum
import copy

class Direction(Enum):
    UP = '^'
    RIGHT = '>'
    DOWN = 'v'
    LEFT = '<'

class Guard:
    def __init__(self, x: int, y: int, direction: Direction):
        self.x = x
        self.y = y
        self.direction = direction

    def turn_right(self):
        directions = [Direction.UP, Direction.RIGHT, Direction.DOWN, Direction.LEFT]
        current_index = directions.index(self.direction)
        self.direction = directions[(current_index + 1) % 4]

    def move_forward(self):
        if self.direction == Direction.UP:
            self.y -= 1
        elif self.direction == Direction.RIGHT:
            self.x += 1
        elif self.direction == Direction.DOWN:
            self.y += 1
        elif self.direction == Direction.LEFT:
            self.x -= 1

def parse_map(grid: List[str]) -> Tuple[List[List[str]], Guard]:
    map_grid = []
    guard = None
    
    for y, row in enumerate(grid):
        map_row = []
        for x, char in enumerate(row):
            if char in ['^', '>', 'v', '<']:
                guard = Guard(x, y, Direction(char))
                map_row.append('.')
            else:
                map_row.append(char)
        map_grid.append(map_row)
    
    return map_grid, guard

def is_in_bounds(x: int, y: int, grid: List[List[str]]) -> bool:
    return 0 <= y < len(grid) and 0 <= x < len(grid[0])

def get_front_position(guard: Guard) -> Tuple[int, int]:
    x, y = guard.x, guard.y
    if guard.direction == Direction.UP:
        y -= 1
    elif guard.direction == Direction.RIGHT:
        x += 1
    elif guard.direction == Direction.DOWN:
        y += 1
    elif guard.direction == Direction.LEFT:
        x -= 1
    return x, y

def simulate_path(grid: List[List[str]], start_guard: Guard) -> Set[Tuple[int, int]]:
    visited = set()
    guard = copy.deepcopy(start_guard)
    
    while True:
        visited.add((guard.x, guard.y))
        
        front_x, front_y = get_front_position(guard)
        
        # Check if guard would move out of bounds
        if not is_in_bounds(front_x, front_y, grid):
            break
            
        # If there's an obstacle in front, turn right
        if grid[front_y][front_x] == '#':
            guard.turn_right()
        else:
            guard.move_forward()
    
    return visited

def find_loop_positions(grid: List[List[str]], start_guard: Guard) -> int:
    original_path = simulate_path(grid, start_guard)
    possible_positions = set()
    
    # Try placing an obstacle at each position
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            # Skip if position is already occupied or is guard's starting position
            if grid[y][x] == '#' or (x == start_guard.x and y == start_guard.y):
                continue
                
            # Create a new grid with the test obstacle
            test_grid = copy.deepcopy(grid)
            test_grid[y][x] = '#'
            
            # Simulate the guard's path with the new obstacle
            test_guard = copy.deepcopy(start_guard)
            visited = set()
            path = []
            current_pos = (test_guard.x, test_guard.y, test_guard.direction.value)
            
            while current_pos not in visited:
                visited.add(current_pos)
                path.append(current_pos)
                
                front_x, front_y = get_front_position(test_guard)
                
                # Check bounds and obstacles
                if not is_in_bounds(front_x, front_y, test_grid) or test_grid[front_y][front_x] == '#':
                    test_guard.turn_right()
                else:
                    test_guard.move_forward()
                    
                current_pos = (test_guard.x, test_guard.y, test_guard.direction.value)
                
                # If we've moved out of bounds, break
                if not is_in_bounds(test_guard.x, test_guard.y, test_grid):
                    break
            
            # If we found a loop (current_pos in visited) and we haven't moved out of bounds
            if current_pos in visited and is_in_bounds(test_guard.x, test_guard.y, test_grid):
                possible_positions.add((x, y))
    
    return len(possible_positions)

def solve_part2(input_lines: List[str]) -> int:
    grid, guard = parse_map(input_lines)
    return find_loop_positions(grid, guard)

# Test with example
example = [
    "....#.....",
    ".........#",
    "..........",
    "..#.......",
    ".......#..",
    "..........",
    ".#..^.....",
    "........#.",
    "#.........",
    "......#..."
]

result = solve_part2(example)
print(f"Example result: {result}")  # Should print 6