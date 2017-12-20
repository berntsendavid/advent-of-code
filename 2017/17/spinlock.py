from collections import deque

buffer = deque([0])
steps = 301
for x in xrange(1, 50000001, 1):
    buffer.rotate(-steps)
    buffer.append(x)

for i, e in enumerate(buffer):
    if e == 0:
        print buffer[(i+1)%len(buffer)]