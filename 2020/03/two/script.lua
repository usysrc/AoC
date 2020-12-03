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

local countTrees = function(tx, ty)
    local position = {x = 0, y = 0}
    local count = 0
    while position.y < j do
        if map[(position.x % width) .. "," .. position.y] == "#" then
            count = count + 1
        end
        position.x = position.x + tx
        position.y = position.y + ty
    end
    return count
end

local trees = {
    countTrees(1, 1), countTrees(3, 1), countTrees(5, 1), countTrees(7, 1),
    countTrees(1, 2)
}

local product = 1
for i, v in ipairs(trees) do
    product = product * v
    print(v)
end
print(product)
