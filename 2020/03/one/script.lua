local map = {}
local width = 0

local j = 0
for line in io.lines() do
    for i = 1, #line do
        map[(i - 1) .. "," .. j] = line:sub(i, i)
        width = i
    end

    j = j + 1
end

local position = {x = 0, y = 0}
local count = 0
while position.y < j do
    if map[(position.x % width) .. "," .. position.y] == "#" then
        count = count + 1
    end
    position.x = position.x + 3
    position.y = position.y + 1
end

print(count)
