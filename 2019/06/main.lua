function string:split(sep)
    local sep, fields = sep or ":", {}
    local pattern = string.format("([^%s]+)", sep)
    self:gsub(pattern, function(c) fields[#fields+1] = c end)
    return fields
 end

local objects = {}
local head 
for line in io.lines("input") do
    local line = line:split(")")
    local left = line[1]
    local right = line[2]
    
    if not objects[left] then
        objects[left] = {
            prev = nil,
            next = {}
        }
    end
    if not objects[right] then
        objects[right] = {
            prev = left,  
            next = {}   
        }
    end
    objects[right].prev = left
    if left == "COM" then
        head = objects[left]
    end
    table.insert(objects[left].next, right)
end

local isInTable = function(t, val)
    for i,v in ipairs(t) do
        if v == val then return true end
    end
    return false
end

local cur = objects[objects["YOU"].prev]
local result = math.huge
local visited = {}

-- backtracking --
function search(steps)
    if visited[cur] then return end
    visited[cur] = true
    if cur.next then
        if isInTable(cur.next, "SAN") then
            result = steps
        else
            for i,v in ipairs(cur.next) do
                local temp = cur
                cur = objects[v]
                search(steps+1)
                cur = temp
            end
        end
    end
    if cur.prev then
        cur = objects[cur.prev]
        search(steps+1)
    end
end
search(0)
print(result)
-- output is 424
