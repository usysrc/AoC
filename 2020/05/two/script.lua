local seats = {}
for line in io.lines() do
    local min, max = 0, 127
    local left, right = 0, 7
    for i=1, #line do
        local c = line:sub(i,i)
        if c == "F" then
            max = max - math.floor((1 + max - min)/2)
        end
        if c == "B" then
            min = min + math.floor((1 + max - min)/2)
        end
        if c == "L" then
            right = right - math.floor((1 + right - left)/2)
        end
        if c == "R" then
            left = left + math.floor((1 + right - left)/2)
        end
    end
    seats[#seats + 1] = {
        row = min,
        column = left,
        id = min * 8 + left
    }
end

local map = {}
for i,v in ipairs(seats) do
    map[v.id] = v
end

for i,v in ipairs(seats) do
    if not map[v.id+1] and map[v.id+2] then
        print(v.id+1, v.row, v.column+1)
    end
end