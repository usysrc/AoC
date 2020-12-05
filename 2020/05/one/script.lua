local maxnum = -1
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
    -- print(min, left)
    -- print()
    local num = min * 8 + left
    if num > maxnum then maxnum = num end
end
print(maxnum)
