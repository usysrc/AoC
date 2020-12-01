intersections = []
map = {}
num = 0

def mark(x,y):
    key = str(x) + "," + str(y)
    if key in map and map[key] != num:
        intersections.append({"x": x, "y": y})
    map[key] = num

f = open("input", "r")    
for line in f:
    num+=1
    x = 0
    y = 0
    instructions = line.split(",")
    for instruction in instructions:
        dir = instruction[0]
        len = instruction[1:]
        if len == '':
            len = 0
        else:
            len = int(len)

        if dir == "R":
            for i in range(x, x + len):
                mark(i,y)
            x += len
        if dir == "L":
            for i in range(x, x - len, -1):
                mark(i,y)
            x -= len
        if dir == "U":
            for i in range(y, y - len, -1):
                mark(x,i)
            y -= len
        if dir == "D":
            for i in range(y, y + len):
                mark(x,i)
            y += len

nearest = 100000000
for intersection in intersections:
    if not(intersection.get('x') == 0 and intersection.get('y') == 0) and abs(intersection.get('x')) + abs(intersection.get('y')) < nearest:
        nearest = abs(intersection.get('x')) + abs(intersection.get('y'))
print(nearest)