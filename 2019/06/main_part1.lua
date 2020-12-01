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

local count = 0
function countNodes(obj)
    if obj.next then
        for i,v in ipairs(obj.next) do
            countNodes(objects[v])
        end
    end
    local cur = obj
    while(cur ~= head) do
        count = count + 1
        cur = objects[cur.prev]
    end
end

countNodes(head)
print(count)
