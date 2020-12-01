intersections = []
map = {}
stepmap = {}
num = 0
step = 0

def mark(x,y):
    key = str(x) + "," + str(y)
    selffound = False
    if key in map:
        for point in map[key]:
            if point.get("number") == num:
                selffound = True
        if selffound == False:
            intersections.append({"x": x, "y": y})
    else:
        map[key] = []
    if selffound == False:
        map[key].append({"number": num, "step": step })

f = open("input", "r")    
for line in f:
    num += 1
    x = 0
    y = 0
    step = 0
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
                step+=1
            x += len
        if dir == "L":
            for i in range(x, x - len, -1):
                mark(i,y)
                step+=1
            x -= len
        if dir == "U":
            for i in range(y, y - len, -1):
                mark(x,i)
                step+=1
            y -= len
        if dir == "D":
            for i in range(y, y + len):
                mark(x,i)
                step+=1
            y += len

nearest = 100000000
for intersection in intersections:
    print(intersection)
    if not(intersection.get('x') == 0 and intersection.get('y') == 0):
        key = str(intersection.get('x')) + "," + str(intersection.get('y'))
        if key in map:
            sum = 0
            for sentry in map[key]:
                if "step" in sentry:
                    sum += sentry.get("step")
            if sum < nearest:
                nearest = sum

print(nearest)
