
local sum = 0
-- read the input file line by line
for line in io.lines("testinput") do
    for left,right in line:gmatch("mul%((%d+),(%d+)%)") do
        sum = sum + left*right
    end
end
print(sum)

function combinedMatches(text, patterns)
    local results = {}
    local positions = {}

    -- Modify text to ensure position matching works
    local searchText = text .. " "

    -- Collect matches for each pattern
    for _, pattern in ipairs(patterns) do
        for match, pos in searchText:gmatch("(" .. pattern .. ")()") do
            table.insert(results, {value = match, pattern = pattern})
            table.insert(positions, tonumber(pos))
            print(match, tonumber(pos))
        end
    end

    -- Sort based on original positions
    local sorted = {}
    for i = 1, #results do
        sorted[i] = {value = results[i].value,
                     pattern = results[i].pattern, 
                     pos = positions[i]}
    end
    table.sort(sorted, function(a, b) return a.pos < b.pos end)

    print("------")
    for i,v in ipairs(sorted) do
        print(v.value)
    end
    
    -- Extract just the values in order
    local orderedMatches = {}
    for _, item in ipairs(sorted) do
        table.insert(orderedMatches, item.value)
    end

    return orderedMatches
end

local sum = 0
local isDo = true
-- read the input file line by line
for line in io.lines("testinput") do
    print(line)
    for _,str in ipairs(combinedMatches(line, {"mul%(%d+,%d+%)", "do%(%)", "don't%(%)"})) do
        if str:match("mul%(%d+,%d+%)") then
            local left,right = str:match("mul%((%d+),(%d+)%)")
            if isDo then
                sum = sum + left*right
            end
        end
        if str:match("do%(%)") then
            isDo = true
        end
        if str:match("don't%(%)") then
            isDo = false
        end
    end
end
print(sum)