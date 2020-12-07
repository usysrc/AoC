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
    for i,v in ipairs(right) do
        if string.match(v, " no ") then break end
        local g = split(v, " ")
        local name = g[3].." "..g[4]
        table.insert(graph[left], 
        {
            name = name,
            amount = tonumber(g[2])
        })
    end
end

local start = graph["shiny gold"]
local value = 0
local go 
go = function(p, mult)
    for i,v in ipairs(p) do
        value = value + v.amount * mult
        go(graph[v.name], mult * v.amount)
    end
end
go(start, 1)
print(value)