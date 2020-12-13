local lines = {}
for line in io.lines() do
    lines[#lines+1] = line
end

local notes = lines[2]

local indices = {}
local buses = {}

local k = 0
for bus in string.gmatch(notes, "%w+") do
    k = k + 1
    if bus == "x" then
    
    else
        buses[k] = bus
        if k ~= 0 then
            indices[#indices+1] = k
        end
    end
end

local time = 0
local step = 1
for _, i in ipairs(indices) do
    local j = 0
    print(i, buses[i])
    while true do
        j = j + 1
        local t = time + j * step
        local bus = buses[i]
        local test = (t+i-1)%bus
        if test == 0 then
            step = step * bus
            time = t
            break
        end
    end
end
print(string.format("%.0f",time ))