def find_steps(filename):
  steps = 0
  maze = []

  file = open(filename, "r")
  for line in file:
    maze.append(int(line.split()[0]))


  length = len(maze)
  index = 0
  while index >= 0 and index < length:
    jump = maze[index]
    if jump >= 3:
      maze[index] -= 1
    else:
      maze[index] += 1
    steps += 1
    index = index + jump

  print steps


find_steps("2017-05.txt")
