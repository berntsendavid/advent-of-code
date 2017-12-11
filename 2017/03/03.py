north = [0, 1]
south = [0, -1]
east = [1, 0]
west = [-1, 0]
currPos = [0, 0]
dirIndex = 0
distance = 0
directions = [south, east, north, west]

def turn(index):
  print 'turn'
  return (index + 1) % 4

def move(currPos, direction):
  x = currPos[0] + direction[0]
  y = currPos[1] + direction[1]

  return [x, y]

index = 0
element = 10

while index < element:
  if index % 2 == 0:
    distance += 1

  dirIndex = turn(dirIndex)
  for d in range(0, distance):
    index += 1
    currPos = move(currPos, directions[dirIndex])
    print currPos

print currPos
