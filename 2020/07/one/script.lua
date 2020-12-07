function split(s, delimiter)
    result = {};
    for match in (s..delimiter):gmatch("(.-)"..delimiter) do
        table.insert(result, match);
    end
    return result;
end

local graph = {}

for line in io.lines() do
    local left, right = table.unpack(split(line, "contain"))
    left = split(left, " ")
    left = left[1].." "..left[2]
    right = split(right, ",")
    graph[left] = graph[left] or {}
    -- print(left)
    for i,v in ipairs(right) do
        if string.match(v, " no ") then break end
        local g = split(v, " ")
        g = g[3].." "..g[4]
        -- print(" ",g)
        table.insert(graph[left], g)
    end
end

local results = {}
local bad = {}

local checkResult
checkResult = function(k)
    print(" ", k)
    if k == "shiny gold" then return true end
    if results[k] then return true end
    if bad[k] then return false end
    local insides = graph[k]
    for i, v in ipairs(insides) do
        if checkResult(v) then 
            results[v] = true
            return true
        else
            bad[v] = true
        end
    end
    bad[k] = true
    return false
end

for k,v in pairs(graph) do
    if checkResult(k) then
        print("====>found")
        results[k] = true
    end
    print("\n")
end

local count = 0
for k,v in pairs(results) do
    if k ~= "shiny gold" then count = count + 1 end
end
print(count)